package services

import (
	"github.com/Danni4421/siresto-be-v2/app/dtos"
	"github.com/Danni4421/siresto-be-v2/app/models"
	"github.com/Danni4421/siresto-be-v2/package/exceptions"
	"gorm.io/gorm"
)

type TransactionService struct {
	DB *gorm.DB
}

type TransactionDetailStore struct {
	Menu     models.Menu
	Quantity int
}

func (t *TransactionService) CreateTransaction(createTransactionDto *dtos.CreateTransactionDTO, userID uint) error {
	return t.DB.Transaction(func(tx *gorm.DB) error {
		transaction := models.Transaction{
			UserID: userID,
			Status: models.TransactionStatusPending,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		menuIDs := make([]int, len(createTransactionDto.Transactions))
		for i, item := range createTransactionDto.Transactions {
			menuIDs[i] = item.MenuID
		}

		var menus []models.Menu
		if err := tx.Where("id IN ?", menuIDs).Find(&menus).Error; err != nil {
			return err
		}

		menuMap := make(map[int64]models.Menu)
		for _, menu := range menus {
			menuMap[menu.ID] = menu
		}

		if len(menuMap) != len(menuIDs) {
			return exceptions.NewBadRequest("Some menus not found")
		}

		var totalPrice float64
		details := make([]models.TransactionDetail, len(createTransactionDto.Transactions))

		for i, item := range createTransactionDto.Transactions {
			menu := menuMap[int64(item.MenuID)]
			itemTotal := menu.Price * float64(item.Quantity)
			totalPrice += itemTotal

			details[i] = models.TransactionDetail{
				TransactionID: transaction.ID,
				MenuID:        item.MenuID,
				Quantity:      item.Quantity,
				TotalPrice:    itemTotal,
			}
		}

		transaction.TotalPrice = int(totalPrice)
		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}

		return tx.Create(&details).Error
	})
}

func (t *TransactionService) GetTransactionByID(transactionID uint, userID uint, isCustomer bool) (any, error) {
	transaction := &models.Transaction{}
	if err := t.DB.Where("id = ? AND user_id = ?", transactionID, userID).Preload("TransactionDetails").First(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *TransactionService) GetAllTransactions(userID uint, isCustomer bool) (any, error) {
	var transactions []models.Transaction

	query := t.DB.Preload("TransactionDetails.Menu").Preload("User")
	if isCustomer {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}

	if !isCustomer {
		return transactions, nil
	}

	result := []dtos.TransactionResponseDTO{}
	for _, transaction := range transactions {
		transactionDto := &dtos.TransactionResponseDTO{
			ID:         transaction.ID,
			TotalPrice: transaction.TotalPrice,
			Status:     transaction.Status,
			CreatedAt:  transaction.CreatedAt,
			UpdatedAt:  transaction.UpdatedAt,
			Details:    make([]dtos.TransactionDetailResponseDTO, len(transaction.TransactionDetails)),
		}
		for i, detail := range transaction.TransactionDetails {
			transactionDto.Details[i] = dtos.TransactionDetailResponseDTO{
				Menu: dtos.MenuResponseDTO{
					ID:          detail.MenuID,
					Name:        detail.Menu.Name,
					Price:       detail.Menu.Price,
					Description: detail.Menu.Description,
				},
				Quantity:   detail.Quantity,
				TotalPrice: detail.TotalPrice,
			}
		}
		result = append(result, *transactionDto)
	}

	return &result, nil
}

func (t *TransactionService) UpdateTransactionStatus(id int, status models.TransactionStatus) error {
	var transaction models.Transaction
	if err := t.DB.First(&transaction, id).Error; err != nil {
		return exceptions.NewNotFound("Transaction not found")
	}

	transaction.Status = status
	if err := t.DB.Save(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (t *TransactionService) DeleteTransaction(id int) error {
	var transaction models.Transaction
	if err := t.DB.First(&transaction, id).Error; err != nil {
		return exceptions.NewNotFound("Transaction not found")
	}

	if err := t.DB.Delete(&transaction).Error; err != nil {
		return err
	}

	return nil
}
