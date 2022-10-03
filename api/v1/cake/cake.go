package cake

import (
	"github.com/adhiana46/cake-store-restfulapi/configs"
	repository "github.com/adhiana46/cake-store-restfulapi/internal/repository/cake"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandlerFiber(r fiber.Router) {
	// create service & repository
	app := configs.GetInstance()

	repo := repository.NewCakeRepositoryMysql(app.DB)
	service := service.NewCakeService(repo)
	// repo := repository.NewUserRepositoryPg(cfg.DB)
	// service := service.NewUserService(repo, cfg.RDB)

	group := r.Group("cakes")

	group.Get("/", getCakeListHandlerFiber(service))
	group.Get("/:id", getCakeDetailHandlerFiber(service))
	group.Post("/", createCakeHandlerFiber(service))
	group.Patch("/:id", updateCakeHandlerFiber(service))
	group.Delete("/:id", deleteCakeHandlerFiber(service))
}

func RegisterHandlerGin(r gin.IRouter) {
	// create service & repository
	app := configs.GetInstance()

	repo := repository.NewCakeRepositoryMysql(app.DB)
	service := service.NewCakeService(repo)
	// repo := repository.NewUserRepositoryPg(cfg.DB)
	// service := service.NewUserService(repo, cfg.RDB)

	group := r.Group("cakes")

	group.GET("/", getCakeListHandlerGin(service))
	group.GET("/:id", getCakeDetailHandlerGin(service))
	group.POST("/", createCakeHandlerGin(service))
	group.PATCH("/:id", updateCakeHandlerGin(service))
	group.DELETE("/:id", deleteCakeHandlerGin(service))
}
