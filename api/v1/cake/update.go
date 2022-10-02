package cake

import (
	"net/http"
	"reflect"

	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
	validatorv10 "github.com/go-playground/validator/v10"
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

		// validate request
		validator := configs.GetInstance().Validator
		if err := validator.Validate.Struct(req); err != nil {
			validationErrs := err.(validatorv10.ValidationErrors)

			errorFields := map[string][]string{}
			for _, e := range validationErrs {
				field, ok := reflect.TypeOf(&req).Elem().FieldByName(e.Field())

				if ok {
					jsonTag := field.Tag.Get("json")
					if jsonTag == "" {
						jsonTag = e.Field()
					}
					errorFields[jsonTag] = append(errorFields[jsonTag], e.Translate(*validator.Trans))
				} else {
					errorFields[e.Field()] = append(errorFields[e.Field()], e.Translate(*validator.Trans))
				}
			}

			return c.Status(http.StatusBadRequest).JSON(utils.JsonError(http.StatusBadRequest, errorFields))
		}

		responseData, httpcode, err := service.Update(req)
		if err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request().URI().String(), err)
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, nil))
		}

		return c.Status(httpcode).JSON(utils.JsonSuccess(httpcode, responseData))
	}
}
