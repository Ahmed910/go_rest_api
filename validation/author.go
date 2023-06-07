package validation

type Author struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=100"`
	Email     string `json:"email" binding:"required,email"`
}
