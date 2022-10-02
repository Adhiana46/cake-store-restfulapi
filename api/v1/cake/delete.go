package cake

import (
	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func deleteCakeHandler(service service.CakeService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		reqId, err := c.ParamsInt("id")
		if err != nil {
			reqId = 0
		}

		req := requests.DeleteCakeRequest{
			ID: reqId,
		}

		_, httpcode, err := service.Delete(req)
		if err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request().URI().String(), err)
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, nil))
		}

		return c.Status(httpcode).JSON(utils.JsonSuccess(httpcode, nil))
	}
}
