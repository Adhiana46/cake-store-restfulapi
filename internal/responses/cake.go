package responses

import (
	"github.com/adhiana46/cake-store-restfulapi/constants"
	"github.com/adhiana46/cake-store-restfulapi/internal/entities"
)

type CakeResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func CakeToResponse(e *entities.Cake) *CakeResponse {
	r := &CakeResponse{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Rating:      e.Rating,
		Image:       e.Image,
		CreatedAt:   e.CreatedAt.Format(constants.DEFAULT_DATETIME_LAYOUT),
		UpdatedAt:   e.UpdatedAt.Format(constants.DEFAULT_DATETIME_LAYOUT),
	}

	return r
}

func CakeToResponseList(eArr []*entities.Cake) (responses []*CakeResponse) {
	for _, e := range eArr {
		responses = append(responses, CakeToResponse(e))
	}

	return
}
