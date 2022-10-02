package utils

import (
	"testing"

	"github.com/adhiana46/cake-store-restfulapi/configs"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/stretchr/testify/assert"
)

// validate valid request
func TestValidateRequest_Valid(t *testing.T) {
	type dummyRequest struct {
		Title       string  `json:"title" validate:"required,min=3,max=100"`
		Description string  `json:"description" validate:"max=255"`
		Rating      float32 `json:"rating" validate:"numeric,gte=0,lte=10"`
		Image       string  `json:"image" validate:"max=255"`
	}

	// create validator
	id := id.New()
	uni := ut.New(id, id)
	trans, _ := uni.GetTranslator("id")
	validate := validator.New()
	id_translations.RegisterDefaultTranslations(validate, trans)

	validator := &configs.Validator{
		Validate: validate,
		Trans:    &trans,
	}

	req := dummyRequest{
		Title:       "Ini valid title",
		Description: "ini valid description",
		Rating:      5.5,
		Image:       "https://valid-image.com/cheese.jpg",
	}

	isValid, validationErrors := ValidateRequest(validator, req)

	assert.True(t, isValid)
	assert.Nil(t, validationErrors)
}

// validate invalid request
func TestValidateRequest_Invalid(t *testing.T) {
	type dummyRequest struct {
		Title       string  `json:"title" validate:"required,min=3,max=100"`
		Description string  `json:"description" validate:"max=255"`
		Rating      float32 `json:"rating" validate:"numeric,gte=0,lte=10"`
		Image       string  `json:"image" validate:"max=255"`
	}

	// create validator
	id := id.New()
	uni := ut.New(id, id)
	trans, _ := uni.GetTranslator("id")
	validate := validator.New()
	id_translations.RegisterDefaultTranslations(validate, trans)

	validator := &configs.Validator{
		Validate: validate,
		Trans:    &trans,
	}

	req := dummyRequest{
		Title:       "",
		Description: "ini valid description",
		Rating:      5.5,
		Image:       "https://valid-image.com/cheese.jpg",
	}

	isValid, validationErrors := ValidateRequest(validator, req)

	assert.False(t, isValid)
	assert.NotNil(t, validationErrors)

	assert.Contains(t, validationErrors, "title")
}
