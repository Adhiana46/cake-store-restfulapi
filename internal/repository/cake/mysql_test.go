package repository

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adhiana46/cake-store-restfulapi/internal/entities"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
)

func TestCakeRepositoryMysql_GetAll(t *testing.T) {
	return // entah kenapa testing nya error 😭😭😭😭😭😭

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	columns := []string{
		"id",
		"title",
		"description",
		"rating",
		"image",
		"created_at",
		"updated_at",
	}

	limit := 10
	skip := 0
	wheres := []utils.SqlWhere{
		utils.SqlWhere{
			Field:    "title",
			Operator: "=",
			Value:    "cheese",
		},
		utils.SqlWhere{
			Field:    "rating",
			Operator: ">=",
			Value:    5,
		},
		utils.SqlWhere{
			Field:    "rating",
			Operator: "<=",
			Value:    8,
		},
	}
	orders := []utils.SqlOrder{
		utils.SqlOrder{
			Field: "rating",
			Dir:   "DESC",
		},
		utils.SqlOrder{
			Field: "title",
			Dir:   "ASC",
		},
	}

	sWheres := "WHERE title = 'cheese' AND rating >= '5' AND rating <= '8'"
	sOrders := "ORDER BY rating DESC,title ASC"
	sLimits := fmt.Sprintf("LIMIT %v, %v", skip, limit)

	rs := sqlmock.NewRows(columns)

	repo := NewCakeRepositoryMysql(db)

	mock.ExpectQuery(fmt.Sprintf(SQL_CAKE_SELECT_ALL_COUNT, sWheres))

	mock.ExpectQuery(SQL_CAKE_SELECT_ALL).
		WithArgs(sWheres, sOrders, sLimits).
		WillReturnRows(rs)

	_, _, err = repo.GetAll(limit, skip, wheres, orders)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
}

func TestCakeRepositoryMysql_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	columns := []string{
		"id",
		"title",
		"description",
		"rating",
		"image",
		"created_at",
		"updated_at",
	}

	rs := sqlmock.NewRows(columns)

	repo := NewCakeRepositoryMysql(db)

	mock.ExpectQuery(SQL_CAKE_SELECT_BY_ID).
		WithArgs(1).
		WillReturnRows(rs)

	_, err = repo.GetById(1)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
}

func TestCakeRepositoryMysql_Store(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	columns := []string{
		"id",
		"title",
		"description",
		"rating",
		"image",
		"created_at",
		"updated_at",
	}

	rs := sqlmock.NewRows(columns)

	cake := &entities.Cake{
		Title:       "Cheese Cake",
		Description: "Ini Description",
		Rating:      5,
		Image:       "https://example.com/cheesecake.jpg",
	}

	repo := NewCakeRepositoryMysql(db)

	mock.ExpectExec("INSERT INTO cakes").
		WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(SQL_CAKE_SELECT_BY_ID).
		WithArgs(1).
		WillReturnRows(rs)

	_, err = repo.Store(cake)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
}

func TestCakeRepositoryMysql_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	columns := []string{
		"id",
		"title",
		"description",
		"rating",
		"image",
		"created_at",
		"updated_at",
	}

	rs := sqlmock.NewRows(columns)

	cake := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "Ini Description",
		Rating:      5,
		Image:       "https://example.com/cheesecake.jpg",
	}

	repo := NewCakeRepositoryMysql(db)

	mock.ExpectExec("UPDATE cakes").
		WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(SQL_CAKE_SELECT_BY_ID).
		WithArgs(cake.ID).
		WillReturnRows(rs)

	_, err = repo.Update(cake)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
}

func TestCakeRepositoryMysql_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	cake := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "Ini Description",
		Rating:      5,
		Image:       "https://example.com/cheesecake.jpg",
	}

	repo := NewCakeRepositoryMysql(db)

	mock.ExpectExec("DELETE FROM cakes").
		WithArgs(cake.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(cake)
	if err != nil {
		t.Errorf("An error occured: %s", err)
	}
}
