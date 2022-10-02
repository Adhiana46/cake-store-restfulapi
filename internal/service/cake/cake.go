package service

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"reflect"

	"github.com/adhiana46/cake-store-restfulapi/internal/entities"
	repository "github.com/adhiana46/cake-store-restfulapi/internal/repository/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	"github.com/adhiana46/cake-store-restfulapi/internal/responses"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
)

type cakeService struct {
	repo repository.CakeRepository
}

func NewCakeService(repo repository.CakeRepository) CakeService {
	return &cakeService{
		repo: repo,
	}
}

func (s *cakeService) GetAll(req requests.GetCakeListRequest) ([]*responses.CakeResponse, *utils.ResponsePagination, int, error) {
	sqlWheres := []utils.SqlWhere{}
	sqlOrders, err := utils.ParseStringToSqlOrder(*req.SortBy)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	if req.S != nil {
		sqlWheres = append(sqlWheres, utils.SqlWhere{
			Field:    "title",
			Operator: "LIKE",
			Value:    fmt.Sprintf("%%%s%%", *req.S),
		})
	}
	if req.RatingMin != nil {
		sqlWheres = append(sqlWheres, utils.SqlWhere{
			Field:    "rating",
			Operator: ">=",
			Value:    *req.RatingMin,
		})
	}
	if req.RatingMax != nil {
		sqlWheres = append(sqlWheres, utils.SqlWhere{
			Field:    "rating",
			Operator: "<=",
			Value:    *req.RatingMax,
		})
	}

	skip := (*req.Page - 1) * *req.Limit
	cakes, totalRows, err := s.repo.GetAll(*req.Limit, skip, sqlWheres, sqlOrders)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	pagination := utils.ResponsePagination{
		Size:        len(cakes),
		Total:       totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(*req.Limit))),
		CurrentPage: int(math.Ceil(float64(skip)/float64(*req.Limit))) + 1,
	}

	return responses.CakeToResponseList(cakes), &pagination, http.StatusOK, nil
}

func (s *cakeService) GetById(req requests.GetCakeRequest) (*responses.CakeResponse, int, error) {
	cake, err := s.repo.GetById(req.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if cake.ID == 0 {
		return nil, http.StatusNotFound, nil
	}

	return responses.CakeToResponse(cake), http.StatusOK, nil
}

func (s *cakeService) Create(req requests.CreateCakeRequest) (*responses.CakeResponse, int, error) {
	cake := &entities.Cake{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	}

	insertedCake, err := s.repo.Store(cake)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return responses.CakeToResponse(insertedCake), http.StatusCreated, nil
}

func (s *cakeService) Update(req requests.UpdateCakeRequest) (*responses.CakeResponse, int, error) {
	cake, err := s.repo.GetById(req.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if cake.ID == 0 {
		return nil, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound))
	}

	// Update value of cake, if value in req is not nil
	reqElem := reflect.ValueOf(&req).Elem()
	reqType := reflect.ValueOf(req).Type()
	cakeElem := reflect.ValueOf(cake).Elem()

	for i := 0; i < reqType.NumField(); i++ {
		fieldName := reqType.Field(i).Name

		if fieldName == "ID" {
			continue
		}

		if cakeElem.FieldByName(fieldName).IsValid() && reqElem.Field(i).Elem().IsValid() {
			cakeElem.FieldByName(fieldName).Set(reflect.ValueOf(reqElem.Field(i).Elem().Interface()))
		}
	}

	updatedCake, err := s.repo.Update(cake)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return responses.CakeToResponse(updatedCake), http.StatusOK, nil
}

func (s *cakeService) Delete(req requests.DeleteCakeRequest) (bool, int, error) {
	cake, err := s.repo.GetById(req.ID)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}

	if cake.ID == 0 {
		return false, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound))
	}

	deleted, err := s.repo.Delete(cake)
	if err != nil || !deleted {
		return false, http.StatusInternalServerError, err
	}

	return deleted, http.StatusOK, nil
}
