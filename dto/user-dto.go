package dto

// UserUpdateDto is used by client when PUT update profile
type UserUpdateDto struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password.omitempty" form:"password.omitempty" validate:"ming:6"`
}

// UserCreateDto is used by client when PUT register a user
type UserCreateDto struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password.omitempty" form:"password.omitempty" validate:"ming:6" binding:"required"`
}
