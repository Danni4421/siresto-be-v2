package controllers

import (
	"strconv"

	"github.com/Danni4421/siresto-be-v2/app/dtos"
	"github.com/Danni4421/siresto-be-v2/app/models"
	"github.com/Danni4421/siresto-be-v2/app/services"
	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"github.com/Danni4421/siresto-be-v2/package/utils"
	"github.com/gofiber/fiber/v2"
)

type MenuCategoryController struct {
	MenuCategoryService *services.MenuCategoryService
	UserService *services.UserService
}

func (m *MenuCategoryController) CreateCategory(c *fiber.Ctx) error {
	createMenuCategoryDTO := new(dtos.MenuCategoryDTO)
	if err := utils.ParseAndValidate(c, createMenuCategoryDTO); err != nil {
		return err
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if err := m.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleManager, models.RoleAdmin}); err != nil {
		return err
	}

	category, err := m.MenuCategoryService.CreateCategory(createMenuCategoryDTO.Name, createMenuCategoryDTO.Description)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu category created successfully",
		"data": fiber.Map{
			"category": category,
		},
	})
}

func (m *MenuCategoryController) GetAllCategories(c *fiber.Ctx) error {
	categories, err := m.MenuCategoryService.GetAllCategories()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu categories retrieved successfully",
		"data": fiber.Map{
			"categories": categories,
		},
	})
}

func (m *MenuCategoryController) GetCategoryByID(c *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}

	category, err := m.MenuCategoryService.GetCategoryByID(int16(categoryID))
	if err != nil {
		return exceptions.NewNotFound("Menu category not found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu category retrieved successfully",
		"data": fiber.Map{
			"category": category,
		},
	})
}

func (m *MenuCategoryController) UpdateCategory(c *fiber.Ctx) error {
	updateCategoryDTO := new(dtos.UpdateMenuCategoryDTO)
	
	if err := utils.ParseAndValidate(c, updateCategoryDTO); err != nil {
		return err
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if err := m.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleManager, models.RoleAdmin}); err != nil {
		return err
	}


	categoryID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}

	updatedCategory, err := m.MenuCategoryService.UpdateCategory(int16(categoryID), updateCategoryDTO)

	if err != nil {
		return exceptions.NewInternalServerError("Failed to update category")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu category updated successfully",
		"data": fiber.Map{
			"category": updatedCategory,
		},
	})
}

func (m *MenuCategoryController) DeleteCategory(c *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if err := m.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleManager, models.RoleAdmin}); err != nil {
		return err
	}


	err = m.MenuCategoryService.DeleteCategory(int16(categoryID))
	if err != nil {
		return exceptions.NewInternalServerError("Failed to delete category")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu category deleted successfully",
	})
}