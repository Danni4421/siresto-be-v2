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

type MenuController struct {
	MenuService *services.MenuService
	MenuCategoryService *services.MenuCategoryService
	UserService *services.UserService
}

func (controller MenuController) CreateMenu(c *fiber.Ctx) error {
	createMenuDTO := new(dtos.CreateMenuDTO)

	if err := utils.ParseAndValidate(c, createMenuDTO); err != nil {
		return err
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if err := controller.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleManager, models.RoleAdmin}); err != nil {
		return err
	}


	if len(createMenuDTO.Categories) == 0 {
		return exceptions.NewBadRequest("At least one category must be provided")
	}

	exists, err := controller.MenuService.VerifyMenuExists(createMenuDTO.Name)
	if err != nil {
		return err
	}
	if exists {
		return exceptions.NewBadRequest("Menu with this name already exists")
	}

	menu, err := controller.MenuService.CreateMenu(createMenuDTO)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu created successfully",
		"data": fiber.Map{
			"menu": menu,
		},
	})
}

func (controller MenuController) GetMenus(c *fiber.Ctx) error {
	menus, err := controller.MenuService.GetMenus()

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Menus fetched successfully",
		"data": fiber.Map{
			"menus": menus,
		},
	})
}

func (controller MenuController) GetMenuByID(c *fiber.Ctx) error {
	menuID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid menu ID")
	}

	menu, err := controller.MenuService.GetMenuByID(uint(menuID))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu fetched successfully",
		"data": fiber.Map{
			"menu": menu,
		},
	})
}

func (controller MenuController) UpdateMenu(c *fiber.Ctx) error {
	updateMenuDTO := new(dtos.UpdateMenuDTO)
	if err := utils.ParseAndValidate(c, updateMenuDTO); err != nil {
		return err
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if err := controller.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleManager, models.RoleAdmin}); err != nil {
		return err
	}

	menuID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid menu ID")
	}

	menu, err := controller.MenuService.UpdateMenu(uint(menuID), updateMenuDTO)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu updated successfully",
		"data": fiber.Map{
			"menu": menu,
		},
	})
}

func (controller MenuController) DeleteMenu(c *fiber.Ctx) error {
	menuID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid menu ID")
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if err := controller.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleManager, models.RoleAdmin}); err != nil {
		return err
	}

	err = controller.MenuService.DeleteMenu(uint(menuID))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu deleted successfully",
	})
}