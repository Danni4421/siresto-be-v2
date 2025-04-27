package dtos

type AuthDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

func (dto *AuthDTO) ErrorMessages() map[string]string {
	return map[string]string{
		"Email.required":    "Email is required",
		"Email.email":       "Email is not a valid email",
		"Password.required": "Password is required",
		"Password.min":      "Password must be at least 8 characters long",
		"Password.max":      "Password must be at most 30 characters long",
	}
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (dto *RefreshTokenDTO) ErrorMessages() map[string]string {
	return map[string]string{
		"RefreshToken.required": "Refresh token is required",
	}
}
