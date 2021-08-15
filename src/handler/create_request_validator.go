package handler

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type CustomValidator interface {
	ValidateStruct(createRequest CreateRequest) []*ValidationErrorResponse
}

type CreateValidator struct {
	validate *validator.Validate
}

// ValidationErrorResponse : defines validation error
type ValidationErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var allowedMap = map[string]string{"jpg": "", "jpeg" :"", "JPG": "", "png": "", "tiff": "", "tif": "", "pdf": ""}

func(cv *CreateValidator) ValidateStruct(createRequest CreateRequest) []*ValidationErrorResponse {
	var errors []*ValidationErrorResponse
	validate := validator.New()
	validate.RegisterValidation("allowed-extensions", validateExtensions)

	err := validate.Struct(createRequest)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func validateExtensions(fl validator.FieldLevel) bool {
	s := strings.Split(fl.Field().String(), ".")
	if len(s) != 2 {
		return false
	}
	ext := s[1]
	if _, ok := allowedMap[ext]; ok {
		return true
	}

	return false
}

func ValidatorFactory() *CreateValidator {
	return &CreateValidator{ validator.New()}
}
