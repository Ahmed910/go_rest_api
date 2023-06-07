package dto

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValidationErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []ErrorMsg  `json:"errors"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
