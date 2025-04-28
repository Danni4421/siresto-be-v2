package dtos

type MenuCategoryDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required"`
}

func (m *MenuCategoryDTO) ErrorMessages() map[string]string {
	return map[string]string{
		"Name.required":        "Name is required",
		"Name.min":             "Name must be at least 3 characters",
		"Name.max":             "Name must be at most 255 characters",
		"Description.required": "Description is required",
	}
}

type UpdateMenuCategoryDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required"`
}

func (m *UpdateMenuCategoryDTO) ErrorMessages() map[string]string {
	return map[string]string{
		"Name.required":        "Name is required",
		"Name.min":             "Name must be at least 3 characters",
		"Name.max":             "Name must be at most 255 characters",
		"Description.required": "Description is required",
	}
}	
