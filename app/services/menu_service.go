package services

import (
	"github.com/Danni4421/siresto-be-v2/app/dtos"
	"github.com/Danni4421/siresto-be-v2/app/models"
	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"gorm.io/gorm"
)

type MenuService struct {
	DB *gorm.DB
}

func (s *MenuService) CreateMenu(createMenuDTO *dtos.CreateMenuDTO) (*models.Menu, error) {
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
 
	menu := &models.Menu{
		Name:        createMenuDTO.Name,
		Description: createMenuDTO.Description,
		Price:       createMenuDTO.Price,
	}
	
	if err := tx.Create(menu).Error; err != nil {
		tx.Rollback()
		return nil, exceptions.NewInternalServerError("Failed to create menu")
	}
	
	if len(createMenuDTO.Categories) > 0 {
		var categories []models.MenuCategory
		if err := tx.Where("id IN ?", createMenuDTO.Categories).Find(&categories).Error; err != nil {
			tx.Rollback()
			return nil, exceptions.NewInternalServerError("Failed to fetch categories")
		}
		
		if len(categories) != len(createMenuDTO.Categories) {
			tx.Rollback()
			return nil, exceptions.NewNotFound("One or more categories not found")
		}
		
		if err := tx.Model(menu).Association("Categories").Replace(categories); err != nil {
			tx.Rollback()
			return nil, exceptions.NewInternalServerError("Failed to associate categories with menu")
		}
	}
	
	if err := tx.Commit().Error; err != nil {
		return nil, exceptions.NewInternalServerError("Failed to commit transaction")
	}
	
	return menu, nil
}


func (s *MenuService) GetMenus() ([]models.Menu, error) {
	var menus []models.Menu
	if err := s.DB.Preload("Categories", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "description")
	}).Find(&menus).Error; err != nil {
		return nil, exceptions.NewInternalServerError("Failed to fetch menus")
	}

	return menus, nil
}

func (s *MenuService) GetMenuByID(id uint) (*models.Menu, error) {
	var menu models.Menu
	if err := s.DB.Preload("Categories").First(&menu, id).Error; err != nil {
		return nil, exceptions.NewNotFound("Menu not found")
	}

	return &menu, nil
}

func (s *MenuService) UpdateMenu(id uint, updateMenuDTO *dtos.UpdateMenuDTO) (*models.Menu, error) {
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	menu := &models.Menu{}
	if err := tx.First(menu, id).Error; err != nil {
		tx.Rollback()
		return nil, exceptions.NewNotFound("Menu not found")
	}

	if err := tx.Model(menu).Updates(models.Menu{
		Name:        updateMenuDTO.Name,
		Description: updateMenuDTO.Description,
		Price:       updateMenuDTO.Price,
	}).Error; err != nil {
		tx.Rollback()
		return nil, exceptions.NewInternalServerError("Failed to update menu")
	}

	if len(updateMenuDTO.Categories) > 0 {
		var categories []models.MenuCategory
		if err := tx.Where("id IN ?", updateMenuDTO.Categories).Find(&categories).Error; err != nil {
			tx.Rollback()
			return nil, exceptions.NewInternalServerError("Failed to fetch categories")
		}

		if len(categories) != len(updateMenuDTO.Categories) {
			tx.Rollback()
			return nil, exceptions.NewNotFound("One or more categories not found")
		}

		if err := tx.Model(menu).Association("Categories").Replace(categories); err != nil {
			tx.Rollback()
			return nil, exceptions.NewInternalServerError("Failed to associate categories with menu")
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exceptions.NewInternalServerError("Failed to commit transaction")
	}

	return menu, nil
}

func (s *MenuService) VerifyMenuExists(name string) (bool, error) {
	var menu models.Menu
	if err := s.DB.Where("name = ?", name).First(&menu).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, exceptions.NewInternalServerError("Failed to verify menu existence")
	}
	return true, nil
}


func (s *MenuService) DeleteMenu(id uint) error {
	if err := s.DB.Delete(&models.Menu{}, id).Error; err != nil {
		return exceptions.NewInternalServerError("Failed to delete menu")
	}

	return nil
}