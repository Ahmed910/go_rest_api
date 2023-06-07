package response

import (
	"errors"
	"gilab.com/pragmaticreviews/golang-gin-poc/dto"
	"github.com/go-playground/validator/v10"
)

func getErrorMsg(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fieldError.Param()
	case "gte":
		return "Should be greater than " + fieldError.Param()
	case "min":
		return "Should not be less than " + fieldError.Param() + "characters"
	case "max":
		return "Should not be greater than " + fieldError.Param() + "characters"
	case "url":
		return "This field Should be url"
	}
	return "Unknown error"

}
func GetErrors(err error) ([]dto.ErrorMsg, string) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		allErrors := make([]dto.ErrorMsg, len(ve))
		for i, fe := range ve {
			allErrors[i] = dto.ErrorMsg{Field: fe.Field(), Message: getErrorMsg(fe)}
		}
		return allErrors, allErrors[0].Message
	}
	return nil, ""
}
