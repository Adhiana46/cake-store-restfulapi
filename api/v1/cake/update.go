package cake

import (
	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func updateCakeHandler(service service.CakeService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		reqId, err := c.ParamsInt("id")
		if err != nil {
			reqId = 0
		}

		req := requests.UpdateCakeRequest{}
		req.ID = reqId

		if err := c.BodyParser(&req); err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request().URI().String(), err)
			return c.Status(fiber.StatusInternalServerError).JSON(utils.JsonError(fiber.StatusInternalServerError, nil))
		}

		responseData, httpcode, validationErrs, err := service.Update(req)
		if err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request().URI().String(), err)
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, nil))
		}

		if validationErrs != nil {
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, validationErrs))
		}

		return c.Status(httpcode).JSON(utils.JsonSuccess(httpcode, responseData))
	}
}
