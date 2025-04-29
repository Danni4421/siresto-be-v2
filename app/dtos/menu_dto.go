package dtos

type CreateMenuDTO struct {
	Name        string  `json:"name" validate:"required,min=3,max=255"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Categories  []int16 `json:"categories"`
}

func (dto *CreateMenuDTO) ErrorMessages() map[string]string {
	return map[string]string{
		"Name.required":        "Name is required",
		"Name.min":             "Name must be at least 3 characters long",
		"Name.max":             "Name must be at most 255 characters long",
		"Description.required": "Description is required",
		"Price.required":       "Price is required",
	}
}

type UpdateMenuDTO struct {
	Name        string  `json:"name" validate:"required,min=3,max=255"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Categories  []int16 `json:"categories"`
}

func (dto *UpdateMenuDTO) ErrorMessages() map[string]string {
	return map[string]string{
		"Name.required":        "Name is required",
		"Name.min":             "Name must be at least 3 characters long",
		"Name.max":             "Name must be at most 255 characters long",
		"Description.required": "Description is required",
		"Price.required":       "Price is required",
	}
}

type MenuResponseDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
