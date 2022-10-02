package main

import (
	"database/sql"
	"fmt"

	"github.com/adhiana46/cake-store-restfulapi/api/v1/cake"
	"github.com/adhiana46/cake-store-restfulapi/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger *logrus.Logger
	DB     *sql.DB
}

func main() {
	app := configs.GetInstance()

	r := fiber.New()

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Cake Store API")
	})

	api := r.Group("/api/v1")
	cake.RegisterHandler(api)

	// TODO: change
	// appHost := os.Getenv("APP_HOST")
	// appPort := os.Getenv("APP_PORT")
	appHost := "0.0.0.0"
	appPort := "8000"

	app.Logger.Infof("Server running on %s at port %s\n", appHost, appPort)
	if err := r.Listen(fmt.Sprintf("%s:%s", appHost, appPort)); err != nil {
		app.Logger.Panicf("Can't start the server, error: %s", err)
	}
}
