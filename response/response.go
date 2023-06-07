package response

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse(status bool, message string, data interface{}) dto.ApiResponse {
	return dto.ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func BuildValidationErrorResponse(status bool, message string, data interface{}, errors []dto.ErrorMsg) dto.ValidationErrorResponse {
	return dto.ValidationErrorResponse{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}
