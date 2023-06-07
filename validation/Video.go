package validation

type Video struct {
	Title       string `json:"title" binding:"required,min=2,max=100"` //validate:"is-cool"
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	Author      Author `json:"author" binding:"required"`
}
