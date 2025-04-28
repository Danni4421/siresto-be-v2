package services

import (
	"github.com/Danni4421/siresto-be-v2/app/dtos"
	"github.com/Danni4421/siresto-be-v2/app/models"
	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"gorm.io/gorm"
)

type MenuCategoryService struct {
	DB *gorm.DB
}

func (s *MenuCategoryService) CreateCategory(name, description string) (*models.MenuCategory, error) {
	category := &models.MenuCategory{
		Name:        name,
		Description: description,
	}

	isCategoryExists := s.DB.Find(&models.MenuCategory{}, "name = ?", name).RowsAffected > 0

	if isCategoryExists {
		return nil, exceptions.NewBadRequest("Category with this name already exists")
	}

	if err := s.DB.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (s *MenuCategoryService) GetAllCategories() ([]models.MenuCategory, error) {
	var categories []models.MenuCategory
	if err := s.DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *MenuCategoryService) GetCategoryByID(id int16) (*models.MenuCategory, error) {
	var category models.MenuCategory
	if err := s.DB.Preload("Menus").First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *MenuCategoryService) UpdateCategory(id int16, updateCategoryDTO *dtos.UpdateMenuCategoryDTO) (*models.MenuCategory, error) {
	if err := s.DB.Model(&models.MenuCategory{}).Where("id = ?", id).Updates(models.MenuCategory{
		Name:        updateCategoryDTO.Name,
		Description: updateCategoryDTO.Description,
	}).Error; err != nil {
		return nil, err
	}

	return s.GetCategoryByID(id)
}

func (s *MenuCategoryService) DeleteCategory(id int16) error {
	if err := s.DB.Delete(&models.MenuCategory{}, id).Error; err != nil {
		return err
	}

	return nil
}