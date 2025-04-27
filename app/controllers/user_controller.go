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

type UserController struct {
	UserService *services.UserService
}

func (controller UserController) Register(c *fiber.Ctx) error {
	createUserDTO := new(dtos.CreateUserDTO)

	if err := utils.ParseAndValidate(c, createUserDTO); err != nil {
		return err
	}

	user, err := controller.UserService.CreateUser(&models.User{
		Name:     createUserDTO.Name,
		Phone:    createUserDTO.Phone,
		Email:    createUserDTO.Email,
		Password: createUserDTO.Password,
	})

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully",
		"data": fiber.Map{
			"user": fiber.Map{
				"id":         user.ID,
				"name":       user.Name,
				"phone":      user.Phone,
				"email":      user.Email,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
			},
		},
	})
}

func (controller UserController) GetUsers(c *fiber.Ctx) error {
	users, err := controller.UserService.FindUsers()

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Users retrieved successfully",
		"data": fiber.Map{
			"users": users,
		},
	})
}

func (controller UserController) GetUserByID(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		return exceptions.NewBadRequest("Invalid user ID")
	}

	user, err := controller.UserService.FindUserByID(uint(userID))

	if err != nil {
		return exceptions.NewNotFound("User not found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User retrieved successfully",
		"data": fiber.Map{
			"user": user,
		},
	})
}

func (controller UserController) UpdateUser(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		return exceptions.NewBadRequest("Invalid user ID")
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if authorizedUserID != uint(userID) {
		return exceptions.NewUnauthorized("You are not authorized to update this user")
	}

	updateUserDTO := new(dtos.UpdateUserDTO)

	if err := utils.ParseAndValidate(c, updateUserDTO); err != nil {
		return err
	}

	if updateUserDTO.Name == "" && updateUserDTO.Phone == "" && updateUserDTO.Address == "" {
		return exceptions.NewBadRequest("Update payload is empty, at least one field must be provided")
	}

	_, err = controller.UserService.UpdateUser(uint(userID), &models.User{
		Name:    updateUserDTO.Name,
		Phone:   updateUserDTO.Phone,
		Address: updateUserDTO.Address,
	})

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User updated successfully",
	})
}

func (controller UserController) DeleteUser(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		return exceptions.NewBadRequest("Invalid user ID")
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if authorizedUserID != uint(userID) {
		return exceptions.NewUnauthorized("You are not authorized to delete this user")
	}

	err = controller.UserService.DeleteUser(uint(userID))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
}
