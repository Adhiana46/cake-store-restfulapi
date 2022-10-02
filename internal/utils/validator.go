package utils

import (
	"reflect"

	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(validatorCfg *configs.Validator, req any) (bool, map[string][]string) {
	if err := validatorCfg.Validate.Struct(req); err != nil {
		validationErrs := err.(validator.ValidationErrors)

		errorFields := map[string][]string{}
		for _, e := range validationErrs {
			field, ok := reflect.TypeOf(req).FieldByName(e.Field())
			// field, ok := reflect.TypeOf(&req).Elem().FieldByName(e.Field())

			if ok {
				jsonTag := field.Tag.Get("json")
				if jsonTag == "" {
					jsonTag = e.Field()
				}
				errorFields[jsonTag] = append(errorFields[jsonTag], e.Translate(*validatorCfg.Trans))
			} else {
				errorFields[e.Field()] = append(errorFields[e.Field()], e.Translate(*validatorCfg.Trans))
			}
		}

		return false, errorFields
	}

	return true, nil
}
