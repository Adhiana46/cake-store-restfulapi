package cake

import (
	"net/http"

	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func getCakeListHandlerFiber(service service.CakeService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		req := requests.GetCakeListRequest{}

		if err := c.QueryParser(&req); err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request().URI().String(), err)
			return c.Status(fiber.StatusInternalServerError).JSON(utils.JsonError(fiber.StatusInternalServerError, nil))
		}

		responseData, pagination, httpcode, validationErrs, err := service.GetAll(req)

		if err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request().URI().String(), err)
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, nil))
		}

		if validationErrs != nil {
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, validationErrs))
		}

		jsonRes := utils.JsonSuccessWithPagination(
			httpcode,
			responseData,
			*pagination,
		)
		return c.Status(httpcode).JSON(jsonRes)
	}
}

func getCakeListHandlerGin(service service.CakeService) func(c *gin.Context) {
	return func(c *gin.Context) {
		req := requests.GetCakeListRequest{}

		if err := c.ShouldBindQuery(&req); err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request.RequestURI, err)
			c.JSON(http.StatusInternalServerError, utils.JsonError(fiber.StatusInternalServerError, nil))
			return
		}

		responseData, pagination, httpcode, validationErrs, err := service.GetAll(req)

		if err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request.RequestURI, err)
			c.JSON(httpcode, utils.JsonError(httpcode, nil))
			return
		}

		if validationErrs != nil {
			c.JSON(httpcode, utils.JsonError(httpcode, validationErrs))
			return
		}

		jsonRes := utils.JsonSuccessWithPagination(
			httpcode,
			responseData,
			*pagination,
		)
		c.JSON(httpcode, jsonRes)
	}
}
