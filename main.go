package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/adhiana46/cake-store-restfulapi/api/v1/cake"
	"github.com/adhiana46/cake-store-restfulapi/configs"
	gincors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger *logrus.Logger
	DB     *sql.DB
}

func main() {
	app := configs.GetInstance()

	newGinApp(app)
}

func newFiberApp(app *configs.Configs) {
	r := fiber.New()

	// middlewares
	r.Use(compress.New())
	r.Use(cors.New())

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Cake Store API")
	})

	api := r.Group("/api/v1")
	cake.RegisterHandlerFiber(api)

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	app.Logger.Infof("Server running on %s at port %s\n", appHost, appPort)
	if err := r.Listen(fmt.Sprintf("%s:%s", appHost, appPort)); err != nil {
		app.Logger.Panicf("Can't start the server, error: %s", err)
	}
}

func newGinApp(app *configs.Configs) {
	r := gin.Default()
	r.Use(gincors.Default())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Cake Store API")
	})

	api := r.Group("/api/v1")
	cake.RegisterHandlerGin(api)

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	app.Logger.Infof("Server running on %s at port %s\n", appHost, appPort)
	if err := r.Run(fmt.Sprintf("%s:%s", appHost, appPort)); err != nil {
		app.Logger.Panicf("Can't start the server, error: %s", err)
	}
}
