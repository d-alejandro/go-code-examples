package repository

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"gorm.io/gorm"
)

func (repository *orderRepository) Create(request *request.OrderStoreRequest) (*models.Order, error) {
	agency := models.Agency{
		Name: request.GetAgencyName(),
	}

	userName := request.GetUserName()
	note := request.GetNote()
	adminNote := request.GetAdminNote()

	order := &models.Order{
		Agency:         agency,
		RentalDate:     request.GetRentalDate(),
		GuestCount:     request.GetGuestCount(),
		TransportCount: request.GetTransportCount(),
		UserName:       &userName,
		Email:          request.GetEmail(),
		Phone:          request.GetPhone(),
		Note:           &note,
		AdminNote:      &adminNote,
	}

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(order)
		if result != nil {
			return result.Error
		}
		return nil
	})

	return order, err
}
