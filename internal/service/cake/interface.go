package service

import (
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	"github.com/adhiana46/cake-store-restfulapi/internal/responses"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
)

type CakeService interface {
	GetAll(req requests.GetCakeListRequest) ([]*responses.CakeResponse, *utils.ResponsePagination, int, map[string][]string, error)
	GetById(req requests.GetCakeRequest) (*responses.CakeResponse, int, map[string][]string, error)
	Create(req requests.CreateCakeRequest) (*responses.CakeResponse, int, map[string][]string, error)
	Update(req requests.UpdateCakeRequest) (*responses.CakeResponse, int, map[string][]string, error)
	Delete(req requests.DeleteCakeRequest) (bool, int, map[string][]string, error)
}
