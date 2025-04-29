package controllers

import (
	"strconv"

	"github.com/Danni4421/siresto-be-v2/app/dtos"
	"github.com/Danni4421/siresto-be-v2/app/models"
	"github.com/Danni4421/siresto-be-v2/app/services"
	"github.com/Danni4421/siresto-be-v2/package/utils"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	UserService        *services.UserService
	TransactionService *services.TransactionService
}

func (controller *TransactionController) CreateTransaction(c *fiber.Ctx) error {
	createTransactionDto := new(dtos.CreateTransactionDTO)
	if err := utils.ParseAndValidate(c, createTransactionDto); err != nil {
		return err
	}

	authorizedUserID := uint(c.Locals("userID").(float64))

	if err := controller.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleCustomer}); err != nil {
		return err
	}

	if err := controller.TransactionService.CreateTransaction(createTransactionDto, authorizedUserID); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Transaction created successfully",
	})
}

func (controller *TransactionController) GetTransactions(c *fiber.Ctx) error {
	authorizedUserID := uint(c.Locals("userID").(float64))

	isCustomer := true
	if err := controller.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleCustomer}); err != nil {
		isCustomer = false
	}

	transactions, err := controller.TransactionService.GetAllTransactions(authorizedUserID, isCustomer)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":       "success",
		"message":      "Transactions retrieved successfully",
		"transactions": transactions,
	})
}

func (controller *TransactionController) GetTransactionByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid transaction ID",
		})
	}
	transactionID := uint(id)
	authorizedUserID := uint(c.Locals("userID").(float64))

	isCustomer := true
	if err := controller.UserService.VerifyUserRole(authorizedUserID, []models.UserRole{models.RoleCustomer}); err != nil {
		isCustomer = false
	}

	transaction, err := controller.TransactionService.GetTransactionByID(transactionID, authorizedUserID, isCustomer)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "success",
		"message":     "Transaction retrieved successfully",
		"transaction": transaction,
	})
}
