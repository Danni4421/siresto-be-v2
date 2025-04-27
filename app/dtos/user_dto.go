package dtos

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Phone    string `json:"phone" validate:"required,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

func (dto *CreateUserDTO) ErrorMessages() map[string]string {
	return map[string]string{
		"Name.required":     "Name is required",
		"Name.min":          "Name must be at least 3 characters long",
		"Name.max":          "Name must be at most 50 characters long",
		"Phone.required":    "Phone is required",
		"Phone.max":         "Phone must be exactly 15 characters long",
		"Email.required":    "Email is required",
		"Email.email":       "Email is not a valid email",
		"Password.required": "Password is required",
		"Password.min":      "Password must be at least 8 characters long",
		"Password.max":      "Password must be at most 30 characters long",
	}
}
