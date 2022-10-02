package responses

import (
	"testing"
	"time"

	"github.com/adhiana46/cake-store-restfulapi/constants"
	"github.com/adhiana46/cake-store-restfulapi/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestCakeToResponse(t *testing.T) {
	e := &entities.Cake{
		ID:          123,
		Title:       "Cheese Cakse",
		Description: "Ini Deskripsi",
		Rating:      7.5,
		Image:       "https://example.com/cheesecake.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	expected := &CakeResponse{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Rating:      e.Rating,
		Image:       e.Image,
		CreatedAt:   e.CreatedAt.Format(constants.DEFAULT_DATETIME_LAYOUT),
		UpdatedAt:   e.UpdatedAt.Format(constants.DEFAULT_DATETIME_LAYOUT),
	}

	result := CakeToResponse(e)

	assert.Equal(t, expected, result)
}

func TestCakeToResponseList_Empty(t *testing.T) {
	cakes := []*entities.Cake{}

	expected := []*CakeResponse{}

	result := CakeToResponseList(cakes)

	assert.Equal(t, expected, result)
}

func TestCakeToResponseList_NotEmpty(t *testing.T) {
	cake1 := &entities.Cake{
		ID:          1,
		Title:       "Cheese Cakse 1",
		Description: "Ini Deskripsi 1",
		Rating:      7.5,
		Image:       "https://example.com/cheesecake-1.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	cake2 := &entities.Cake{
		ID:          2,
		Title:       "Cheese Cakse 2",
		Description: "Ini Deskripsi 2",
		Rating:      7.8,
		Image:       "https://example.com/cheesecake-2.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	cakes := []*entities.Cake{
		cake1,
		cake2,
	}

	expected := []*CakeResponse{
		&CakeResponse{
			ID:          cake1.ID,
			Title:       cake1.Title,
			Description: cake1.Description,
			Rating:      cake1.Rating,
			Image:       cake1.Image,
			CreatedAt:   cake1.CreatedAt.Format(constants.DEFAULT_DATETIME_LAYOUT),
			UpdatedAt:   cake1.UpdatedAt.Format(constants.DEFAULT_DATETIME_LAYOUT),
		},
		&CakeResponse{
			ID:          cake2.ID,
			Title:       cake2.Title,
			Description: cake2.Description,
			Rating:      cake2.Rating,
			Image:       cake2.Image,
			CreatedAt:   cake2.CreatedAt.Format(constants.DEFAULT_DATETIME_LAYOUT),
			UpdatedAt:   cake2.UpdatedAt.Format(constants.DEFAULT_DATETIME_LAYOUT),
		},
	}

	result := CakeToResponseList(cakes)

	assert.Equal(t, expected, result)
}
