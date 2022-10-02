package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/adhiana46/cake-store-restfulapi/constants"
	"github.com/adhiana46/cake-store-restfulapi/internal/entities"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
	"github.com/sirupsen/logrus"
)

const (
	SQL_CAKE_SELECT_ALL       = "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes %s %s %s" // replace ? with where,order,limit, dll
	SQL_CAKE_SELECT_ALL_COUNT = "SELECT COUNT(id) AS numrows FROM cakes %s"
	SQL_CAKE_SELECT_BY_ID     = "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?"
	SQL_CAKE_CREATE           = "INSERT INTO cakes(title, description, rating, image) VALUES(?, ?, ?, ?)"
	SQL_CAKE_UPDATE           = "UPDATE cakes SET title=?, description=?, rating=?, image=? WHERE id = ?"
	SQL_CAKE_DELETE           = "DELETE FROM cakes WHERE id = ?"
)

type cakeRepositoryMysql struct {
	db *sql.DB
}

func NewCakeRepositoryMysql(db *sql.DB) CakeRepository {
	return &cakeRepositoryMysql{
		db: db,
	}
}

func (r *cakeRepositoryMysql) GetAll(limit int, skip int, wheres []utils.SqlWhere, orders []utils.SqlOrder) ([]*entities.Cake, int, error) {
	var sWheres, sOrders string = " ", ""

	aWheres := []string{}
	for _, where := range wheres {
		aWheres = append(aWheres, fmt.Sprintf("%s %s '%v'", where.Field, where.Operator, where.Value))
	}
	if len(aWheres) > 0 {
		sWheres = "WHERE " + strings.Join(aWheres, " AND ")
	}

	aOrders := []string{}
	for _, order := range orders {
		aOrders = append(aOrders, fmt.Sprintf("%s %s", order.Field, order.Dir))
	}
	if len(aOrders) > 0 {
		sOrders = "ORDER BY " + strings.Join(aOrders, ",")
	}

	sLimit := fmt.Sprintf("LIMIT %v, %v", skip, limit)

	fmt.Println(sWheres)

	countRow := r.db.QueryRow(fmt.Sprintf(SQL_CAKE_SELECT_ALL_COUNT, sWheres))
	var totalRows int
	err := countRow.Scan(&totalRows)
	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("repository.cakeRepositoryMysql.GetAll: %s", err)
		return nil, 0, err
	}

	rows, err := r.db.Query(fmt.Sprintf(SQL_CAKE_SELECT_ALL, sWheres, sOrders, sLimit))
	if err != nil {
		logrus.Errorf("repository.cakeRepositoryMysql.GetAll: %s", err)
		return nil, 0, err
	}
	defer rows.Close()

	cakes := []*entities.Cake{}
	for rows.Next() {
		cake := entities.Cake{}
		var createdAt, updatedAt string
		err = rows.Scan(
			&cake.ID,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			logrus.Errorf("repository.cakeRepositoryMysql.GetAll: %s", err)
			return nil, 0, err
		}

		if cake.CreatedAt, err = time.Parse(constants.DEFAULT_DATETIME_LAYOUT, createdAt); err != nil {
			logrus.Errorf("repository.cakeRepositoryMysql.GetAll: %s", err)
			return nil, 0, nil
		}
		if cake.UpdatedAt, err = time.Parse(constants.DEFAULT_DATETIME_LAYOUT, updatedAt); err != nil {
			logrus.Errorf("repository.cakeRepositoryMysql.GetAll: %s", err)
			return nil, 0, nil
		}

		cakes = append(cakes, &cake)
	}

	return cakes, totalRows, nil
}

func (r *cakeRepositoryMysql) GetById(id int) (*entities.Cake, error) {
	cake := &entities.Cake{}
	row := r.db.QueryRow(SQL_CAKE_SELECT_BY_ID, id)
	var createdAt, updatedAt string
	err := row.Scan(
		&cake.ID,
		&cake.Title,
		&cake.Description,
		&cake.Rating,
		&cake.Image,
		&createdAt,
		&updatedAt,
	)

	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("repository.cakeRepositoryMysql.GetById: %s", err)
		return cake, err
	}

	if err == nil && cake.ID != 0 {
		if cake.CreatedAt, err = time.Parse(constants.DEFAULT_DATETIME_LAYOUT, createdAt); err != nil {
			logrus.Errorf("repository.cakeRepositoryMysql.GetById: %s", err)
			return cake, err
		}
		if cake.UpdatedAt, err = time.Parse(constants.DEFAULT_DATETIME_LAYOUT, updatedAt); err != nil {
			logrus.Errorf("repository.cakeRepositoryMysql.GetById: %s", err)
			return cake, err
		}
	}

	return cake, nil
}

func (r *cakeRepositoryMysql) Store(cake *entities.Cake) (*entities.Cake, error) {
	res, err := r.db.Exec(SQL_CAKE_CREATE, cake.Title, cake.Description, cake.Rating, cake.Image)
	if err != nil {
		logrus.Errorf("repository.cakeRepositoryMysql.Store: %s", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		logrus.Errorf("repository.cakeRepositoryMysql.Store: %s", err)
		return nil, err
	}

	return r.GetById(int(id))
}

func (r *cakeRepositoryMysql) Update(cake *entities.Cake) (*entities.Cake, error) {
	_, err := r.db.Exec(SQL_CAKE_UPDATE, cake.Title, cake.Description, cake.Rating, cake.Image, cake.ID)
	if err != nil {
		logrus.Errorf("repository.cakeRepositoryMysql.Update: %s", err)
		return nil, err
	}

	return r.GetById(cake.ID)
}

func (r *cakeRepositoryMysql) Delete(cake *entities.Cake) (bool, error) {
	_, err := r.db.Exec(SQL_CAKE_DELETE, cake.ID)
	if err != nil {
		logrus.Errorf("repository.cakeRepositoryMysql.Delete: %s", err)
		return false, err
	}

	return true, nil
}
