package cake

import (
	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func getCakeListHandler(service service.CakeService) func(c *fiber.Ctx) error {
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
