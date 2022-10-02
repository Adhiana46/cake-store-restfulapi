package cake

import (
	"github.com/adhiana46/cake-store-restfulapi/configs"
	repository "github.com/adhiana46/cake-store-restfulapi/internal/repository/cake"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(r fiber.Router) {
	// create service & repository
	app := configs.GetInstance()

	repo := repository.NewCakeRepositoryMysql(app.DB)
	service := service.NewCakeService(repo)
	// repo := repository.NewUserRepositoryPg(cfg.DB)
	// service := service.NewUserService(repo, cfg.RDB)

	group := r.Group("cakes")

	group.Get("/", getCakeListHandler(service))
	group.Get("/:id", getCakeDetailHandler(service))
	group.Post("/", createCakeHandler(service))
	group.Patch("/:id", updateCakeHandler(service))
	group.Delete("/:id", deleteCakeHandler(service))
}
