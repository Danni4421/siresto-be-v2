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

type CreateInternalUserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Phone    string `json:"phone" validate:"required,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
	Role     string `json:"role" validate:"required,oneof=manager"`
}

func (dto *CreateInternalUserDTO) ErrorMessages() map[string]string {
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
		"Role.required":     "Role is required",
		"Role.oneof":        "Role must be manager",
	}
}

type CreateAdminUserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Phone    string `json:"phone" validate:"required,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
	AppPassword string `json:"app_password" validate:"required"`
}

func (dto *CreateAdminUserDTO) ErrorMessages() map[string]string {
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
		"AppPassword.required": "App password is required",
	}
}

type UpdateUserDTO struct {
	Name    string `json:"name" validate:"omitempty,min=3,max=50"`
	Phone   string `json:"phone" validate:"omitempty,max=15"`
	Address string `json:"address" validate:"omitempty,max=255"`
}

func (dto *UpdateUserDTO) ErrorMessages() map[string]string {
	return map[string]string{
		"Name.min":    "Name must be at least 3 characters long",
		"Name.max":    "Name must be at most 50 characters long",
		"Phone.max":   "Phone must be exactly 15 characters long",
		"Address.max": "Address must be at most 255 characters long",
	}
}
