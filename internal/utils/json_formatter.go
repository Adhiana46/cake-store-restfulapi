package utils

import (
	"net/http"
)

type Response struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Errors     any    `json:"errors,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
}

type ResponsePagination struct {
	Size        int `json:"size"`
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
}

func JsonSuccess(status int, data any) Response {
	return Response{
		Status:  status,
		Message: http.StatusText(status),
		Data:    data,
	}
}

func JsonSuccessWithPagination(status int, data any, pagination ResponsePagination) Response {
	return Response{
		Status:     status,
		Message:    http.StatusText(status),
		Data:       data,
		Pagination: pagination,
	}
}

func JsonError(status int, errs any) Response {
	return Response{
		Status:  status,
		Message: http.StatusText(status),
		Errors:  errs,
	}
}
