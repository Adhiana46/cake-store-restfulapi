package cake

import (
	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func deleteCakeHandlerFiber(service service.CakeService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		reqId, err := c.ParamsInt("id")
		if err != nil {
			reqId = 0
		}

		req := requests.DeleteCakeRequest{
			ID: reqId,
		}

		_, httpcode, validationErrs, err := service.Delete(req)
		if err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request().URI().String(), err)
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, nil))
		}

		if validationErrs != nil {
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, validationErrs))
		}

		return c.Status(httpcode).JSON(utils.JsonSuccess(httpcode, nil))
	}
}

func deleteCakeHandlerGin(service service.CakeService) func(c *gin.Context) {
	return func(c *gin.Context) {
		req := requests.DeleteCakeRequest{}

		if err := c.ShouldBindUri(&req); err != nil {
			req.ID = 0
		}

		_, httpcode, validationErrs, err := service.Delete(req)
		if err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request.RequestURI, err)
			c.JSON(httpcode, utils.JsonError(httpcode, nil))
			return
		}

		if validationErrs != nil {
			c.JSON(httpcode, utils.JsonError(httpcode, validationErrs))
			return
		}

		c.JSON(httpcode, utils.JsonSuccess(httpcode, nil))
	}
}
