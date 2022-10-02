package repository

import (
	"github.com/adhiana46/cake-store-restfulapi/internal/entities"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
)

type CakeRepository interface {
	// /cakes?limit=10&skip=0&s=bla bla&rating_min=2.5&rating_max=4.5&sort_by=rating.desc,title.asc
	GetAll(limit int, skip int, wheres []utils.SqlWhere, orders []utils.SqlOrder) ([]*entities.Cake, int, error)
	GetById(id int) (*entities.Cake, error)
	Store(cake *entities.Cake) (*entities.Cake, error)
	Update(cake *entities.Cake) (*entities.Cake, error)
	Delete(cake *entities.Cake) (bool, error)
}
