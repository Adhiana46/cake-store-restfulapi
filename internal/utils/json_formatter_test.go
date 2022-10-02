package utils

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonSuccess(t *testing.T) {
	dummyData := map[string]string{
		"nama":   "Adhiana Mastur",
		"alamat": "Yogyakarta",
	}
	expectedResult := Response{
		Status:  200,
		Message: http.StatusText(200),
		Data:    dummyData,
	}

	assert.Equal(t, expectedResult, JsonSuccess(200, dummyData))
}

func TestJsonSuccessWithPagination(t *testing.T) {
	dummyData := map[string]string{
		"nama":   "Adhiana Mastur",
		"alamat": "Yogyakarta",
	}
	pagination := ResponsePagination{
		Size:        10,
		Total:       100,
		TotalPages:  10,
		CurrentPage: 1,
	}
	expectedResult := Response{
		Status:     200,
		Message:    http.StatusText(200),
		Data:       dummyData,
		Pagination: pagination,
	}

	assert.Equal(t, expectedResult, JsonSuccessWithPagination(200, dummyData, pagination))
}

func TestJsonError(t *testing.T) {
	err := errors.New("Ini Error")
	expectedResult := Response{
		Status:  500,
		Message: http.StatusText(500),
		Errors:  err,
	}

	assert.Equal(t, expectedResult, JsonError(500, err))
}
