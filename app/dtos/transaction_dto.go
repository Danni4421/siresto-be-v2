package dtos

import (
	"time"

	"github.com/Danni4421/siresto-be-v2/app/models"
)

type TransactionDTO struct {
	MenuID   int `json:"menu_id"`
	Quantity int `json:"quantity"`
}

type CreateTransactionDTO struct {
	Transactions []TransactionDTO `json:"transactions" validate:"required,min=1"`
}

type TransactionResponseDTO struct {
	ID         int                            `json:"id"`
	TotalPrice int                            `json:"total_price"`
	Status     models.TransactionStatus       `json:"status"`
	CreatedAt  time.Time                      `json:"created_at"`
	UpdatedAt  time.Time                      `json:"updated_at"`
	DeletedAt  time.Time                      `json:"deleted_at"`
	Details    []TransactionDetailResponseDTO `json:"transaction_detail"`
}

type TransactionDetailResponseDTO struct {
	Menu       MenuResponseDTO `json:"menu"`
	Quantity   int             `json:"quantity"`
	TotalPrice float64         `json:"total_price"`
}
