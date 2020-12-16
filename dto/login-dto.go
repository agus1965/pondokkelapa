package dto

// LoginDto is a model that used when client post form /login url
type LoginDto struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password.omitempty" form:"password.omitempty" validate:"ming:6"`
}
	