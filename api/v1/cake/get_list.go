package cake

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/adhiana46/cake-store-restfulapi/internal/requests"
	service "github.com/adhiana46/cake-store-restfulapi/internal/service/cake"
	"github.com/adhiana46/cake-store-restfulapi/internal/utils"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func getCakeListHandler(service service.CakeService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		req := requests.GetCakeListRequest{}

		if err := c.QueryParser(&req); err != nil {
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

		// jika limit 0 / tidak user request tdk input limit, default 30
		if req.Limit == nil {
			req.Limit = new(int)
			*req.Limit = 30
		}
		if req.Page == nil {
			req.Page = new(int)
			*req.Page = 1
		}
		if req.SortBy == nil {
			req.SortBy = new(string)
			*req.SortBy = "rating.desc,title.asc" // default sorting
		}

		fmt.Println(req.RatingMin, req.RatingMax)

		responseData, pagination, httpcode, err := service.GetAll(req)

		if err != nil {
			configs.GetInstance().Logger.Errorf("%s: %s", c.Request().URI().String(), err)
			return c.Status(httpcode).JSON(utils.JsonError(httpcode, nil))
		}

		jsonRes := utils.JsonSuccessWithPagination(
			httpcode,
			responseData,
			*pagination,
		)
		return c.Status(httpcode).JSON(jsonRes)
	}
}
