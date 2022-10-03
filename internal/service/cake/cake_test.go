package service

import (
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adhiana46/cake-store-restfulapi/constants"
	"github.com/adhiana46/cake-store-restfulapi/internal/entities"
	repository "github.com/adhiana46/cake-store-restfulapi/internal/repository/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	"github.com/adhiana46/cake-store-restfulapi/internal/responses"
	"github.com/stretchr/testify/assert"
)

func TestGetAll_ValidationError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	req := requests.GetCakeListRequest{}
	limit := 128 // max 100
	page := 0    // min 1
	s := "Ullamco excepteur nisi sint incididunt Lorem ea. Ullamco excepteur nisi sint incididunt Lorem ea. Ullamco excepteur nisi sint incididunt Lorem ea."
	ratingMin := float32(-1)
	ratingMax := float32(100)

	req.Limit = &limit
	req.Page = &page
	req.S = &s
	req.RatingMin = &ratingMin
	req.RatingMax = &ratingMax

	response, pagination, httpcode, validationErrs, err := service.GetAll(req)

	assert.Nil(t, response)
	assert.Nil(t, pagination)
	assert.Equal(t, http.StatusBadRequest, httpcode)
	assert.NotNil(t, validationErrs)
	assert.Nil(t, err)

	assert.Contains(t, validationErrs, "limit")
	assert.Contains(t, validationErrs, "page")
	assert.Contains(t, validationErrs, "s")
	assert.Contains(t, validationErrs, "rating_min")
	assert.Contains(t, validationErrs, "rating_max")
}

func TestGetAll_ValidationErrorInvalidSortBy(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	req := requests.GetCakeListRequest{}
	sortby := "invalid_format"

	req.SortBy = &sortby

	response, pagination, httpcode, validationErrs, err := service.GetAll(req)

	assert.Nil(t, response)
	assert.Nil(t, pagination)
	assert.Equal(t, http.StatusBadRequest, httpcode)
	assert.Nil(t, validationErrs)
	assert.NotNil(t, err)
}

func TestGetById_Found(t *testing.T) {
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

	rs := sqlmock.NewRows(columns).
		FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00")

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(1).
		WillReturnRows(rs)

	req := requests.GetCakeRequest{
		ID: 1,
	}
	response, httpcode, validationErrs, err := service.GetById(req)

	expectedResponse := &responses.CakeResponse{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   "2022-10-02 20:40:00",
		UpdatedAt:   "2022-10-02 20:40:00",
	}
	expectedHttpcode := http.StatusOK

	assert.Equal(t, expectedResponse, response)
	assert.Equal(t, expectedHttpcode, httpcode)
	assert.Nil(t, validationErrs)
	assert.Nil(t, err)
}

func TestGetById_NotFound(t *testing.T) {
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

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(2).
		WillReturnRows(rs)

	req := requests.GetCakeRequest{
		ID: 2,
	}
	response, httpcode, validationErrs, err := service.GetById(req)

	assert.Nil(t, response)
	assert.Equal(t, http.StatusNotFound, httpcode)
	assert.Nil(t, validationErrs)
	assert.Nil(t, err)
}

func TestGetById_ValidationError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	// columns := []string{
	// 	"id",
	// 	"title",
	// 	"description",
	// 	"rating",
	// 	"image",
	// 	"created_at",
	// 	"updated_at",
	// }

	// rs := sqlmock.NewRows(columns)

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	// mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
	// 	WithArgs(2).
	// 	WillReturnRows(rs)

	req := requests.GetCakeRequest{}
	response, httpcode, validationErrs, err := service.GetById(req)

	assert.Nil(t, response)
	assert.Equal(t, http.StatusBadRequest, httpcode)
	assert.NotNil(t, validationErrs)
	assert.Nil(t, err)
}

func TestCreate_Success(t *testing.T) {
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

	rs := sqlmock.NewRows(columns).
		FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00")

	cake := &entities.Cake{
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	mock.ExpectExec("INSERT INTO cakes").
		WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(1).
		WillReturnRows(rs)

	req := requests.CreateCakeRequest{
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
	}

	response, httpcode, validationErrs, err := service.Create(req)

	expectedResponse := &responses.CakeResponse{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   "2022-10-02 20:40:00",
		UpdatedAt:   "2022-10-02 20:40:00",
	}
	expectedHttpcode := http.StatusCreated

	assert.Equal(t, expectedResponse, response)
	assert.Equal(t, expectedHttpcode, httpcode)
	assert.Nil(t, validationErrs)
	assert.Nil(t, err)
}

func TestCreate_ValidationError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	// columns := []string{
	// 	"id",
	// 	"title",
	// 	"description",
	// 	"rating",
	// 	"image",
	// 	"created_at",
	// 	"updated_at",
	// }

	// rs := sqlmock.NewRows(columns)

	cake := &entities.Cake{
		Title:       "",
		Description: "ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi ini deskripsi",
		Rating:      17.5,
		Image:       "https://image.com/cheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheese.jpg",
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	// mock.ExpectExec("INSERT INTO cakes").
	// 	WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image).
	// 	WillReturnResult(sqlmock.NewResult(1, 1))

	// mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
	// 	WithArgs(1).
	// 	WillReturnRows(rs)

	req := requests.CreateCakeRequest{
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
	}

	response, httpcode, validationErrs, err := service.Create(req)

	assert.Nil(t, response)
	assert.Equal(t, http.StatusBadRequest, httpcode)
	assert.NotNil(t, validationErrs)
	assert.Nil(t, err)

	assert.Contains(t, validationErrs, "title")
	assert.Contains(t, validationErrs, "description")
	assert.Contains(t, validationErrs, "rating")
	assert.Contains(t, validationErrs, "image")
}

func TestUpdate_Success(t *testing.T) {
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

	rs := sqlmock.NewRows(columns).
		FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00").
		FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00")

	createdAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")
	updatedAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")

	cake := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(cake.ID).
		WillReturnRows(rs)

	mock.ExpectExec("UPDATE cakes").
		WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(cake.ID).
		WillReturnRows(rs)

	req := requests.UpdateCakeRequest{
		ID:          cake.ID,
		Title:       &cake.Title,
		Description: &cake.Description,
		Rating:      &cake.Rating,
		Image:       &cake.Image,
	}

	response, httpcode, validationErrs, err := service.Update(req)

	expectedResponse := &responses.CakeResponse{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   "2022-10-02 20:40:00",
		UpdatedAt:   "2022-10-02 20:40:00",
	}
	expectedHttpcode := http.StatusOK

	assert.Equal(t, expectedResponse, response)
	assert.Equal(t, expectedHttpcode, httpcode)
	assert.Nil(t, validationErrs)
	assert.Nil(t, err)
}

func TestUpdate_NotFound(t *testing.T) {
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

	createdAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")
	updatedAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")

	cake := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(cake.ID).
		WillReturnRows(rs)

	req := requests.UpdateCakeRequest{
		ID:          cake.ID,
		Title:       &cake.Title,
		Description: &cake.Description,
		Rating:      &cake.Rating,
		Image:       &cake.Image,
	}

	response, httpcode, validationErrs, err := service.Update(req)

	assert.Nil(t, response)
	assert.Equal(t, http.StatusNotFound, httpcode)
	assert.Nil(t, validationErrs)
	assert.NotNil(t, err)
}

func TestUpdatePartial_Success(t *testing.T) {
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

	rs := sqlmock.NewRows(columns).
		FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00").
		FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00")

	createdAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")
	updatedAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")

	cake := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(cake.ID).
		WillReturnRows(rs)

	mock.ExpectExec("UPDATE cakes").
		WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(cake.ID).
		WillReturnRows(rs)

	req := requests.UpdateCakeRequest{
		ID:    cake.ID,
		Image: &cake.Image,
	}

	response, httpcode, validationErrs, err := service.Update(req)

	expectedResponse := &responses.CakeResponse{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   "2022-10-02 20:40:00",
		UpdatedAt:   "2022-10-02 20:40:00",
	}
	expectedHttpcode := http.StatusOK

	assert.Equal(t, expectedResponse, response)
	assert.Equal(t, expectedHttpcode, httpcode)
	assert.Nil(t, validationErrs)
	assert.Nil(t, err)
}

func TestUpdate_ValidationError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	// columns := []string{
	// 	"id",
	// 	"title",
	// 	"description",
	// 	"rating",
	// 	"image",
	// 	"created_at",
	// 	"updated_at",
	// }

	// rs := sqlmock.NewRows(columns).
	// 	FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00").
	// 	FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00")

	createdAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")
	updatedAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")

	cake := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	// mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
	// 	WithArgs(cake.ID).
	// 	WillReturnRows(rs)

	// mock.ExpectExec("UPDATE cakes").
	// 	WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.ID).
	// 	WillReturnResult(sqlmock.NewResult(1, 1))

	// mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
	// 	WithArgs(cake.ID).
	// 	WillReturnRows(rs)

	title := "ab"
	description := "Enim aute irure consectetur incididunt dolore ad sit.Enim aute irure consectetur incididunt dolore ad sit.Enim aute irure consectetur incididunt dolore ad sit.Enim aute irure consectetur incididunt dolore ad sit.Enim aute irure consectetur incididunt dolore ad sit."
	rating := float32(11.5)
	image := "https://image.com/cheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheesecheese.jpg"

	req := requests.UpdateCakeRequest{
		ID:          cake.ID,
		Title:       &title,
		Description: &description,
		Rating:      &rating,
		Image:       &image,
	}

	response, httpcode, validationErrs, err := service.Update(req)

	assert.Nil(t, response)
	assert.Equal(t, http.StatusBadRequest, httpcode)
	assert.NotNil(t, validationErrs)
	assert.Nil(t, err)

	assert.Contains(t, validationErrs, "title")
	assert.Contains(t, validationErrs, "description")
	assert.Contains(t, validationErrs, "rating")
	assert.Contains(t, validationErrs, "image")
}

func TestUpdatePartial_ValidationError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	// columns := []string{
	// 	"id",
	// 	"title",
	// 	"description",
	// 	"rating",
	// 	"image",
	// 	"created_at",
	// 	"updated_at",
	// }

	// rs := sqlmock.NewRows(columns).
	// 	FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00").
	// 	FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00")

	createdAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")
	updatedAt, _ := time.Parse(constants.DEFAULT_DATETIME_LAYOUT, "2022-10-02 20:40:00")

	cake := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	// mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
	// 	WithArgs(cake.ID).
	// 	WillReturnRows(rs)

	// mock.ExpectExec("UPDATE cakes").
	// 	WithArgs(cake.Title, cake.Description, cake.Rating, cake.Image, cake.ID).
	// 	WillReturnResult(sqlmock.NewResult(1, 1))

	// mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
	// 	WithArgs(cake.ID).
	// 	WillReturnRows(rs)

	title := "ab"

	req := requests.UpdateCakeRequest{
		ID:    cake.ID,
		Title: &title,
	}

	response, httpcode, validationErrs, err := service.Update(req)

	assert.Nil(t, response)
	assert.Equal(t, http.StatusBadRequest, httpcode)
	assert.NotNil(t, validationErrs)
	assert.Nil(t, err)

	assert.Contains(t, validationErrs, "title")
	assert.NotContains(t, validationErrs, "description")
	assert.NotContains(t, validationErrs, "rating")
	assert.NotContains(t, validationErrs, "image")
}

func TestDelete_Success(t *testing.T) {
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

	rs := sqlmock.NewRows(columns).
		FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00")

	cake := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cake",
		Description: "deskripsi",
		Rating:      7.5,
		Image:       "https://image.com/cheese.jpg",
	}

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(cake.ID).
		WillReturnRows(rs)

	mock.ExpectExec("DELETE FROM cakes").
		WithArgs(cake.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := requests.DeleteCakeRequest{
		ID: cake.ID,
	}

	response, httpcode, validationErrs, err := service.Delete(req)

	assert.True(t, response)
	assert.Equal(t, http.StatusOK, httpcode)
	assert.Nil(t, validationErrs)
	assert.Nil(t, err)
}

func TestDelete_NotFound(t *testing.T) {
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

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	mock.ExpectQuery(repository.SQL_CAKE_SELECT_BY_ID).
		WithArgs(2).
		WillReturnRows(rs)

	req := requests.DeleteCakeRequest{
		ID: 2,
	}

	response, httpcode, validationErrs, err := service.Delete(req)

	assert.False(t, response)
	assert.Equal(t, http.StatusNotFound, httpcode)
	assert.Nil(t, validationErrs)
	assert.NotNil(t, err)
}

func TestDelete_ValidationError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when creating mock db", err)
	}

	// columns := []string{
	// 	"id",
	// 	"title",
	// 	"description",
	// 	"rating",
	// 	"image",
	// 	"created_at",
	// 	"updated_at",
	// }

	// rs := sqlmock.NewRows(columns).
	// 	FromCSVString("1,Cheese Cake,deskripsi,7.5,https://image.com/cheese.jpg,2022-10-02 20:40:00,2022-10-02 20:40:00")

	repo := repository.NewCakeRepositoryMysql(db)
	service := NewCakeService(repo)

	req := requests.DeleteCakeRequest{}

	response, httpcode, validationErrs, err := service.Delete(req)

	assert.False(t, response)
	assert.Equal(t, http.StatusBadRequest, httpcode)
	assert.NotNil(t, validationErrs)
	assert.Nil(t, err)

	assert.Contains(t, validationErrs, "ID")
}
