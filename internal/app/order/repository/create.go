package repository

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"gorm.io/gorm"
)

func (repository *orderRepository) Create(req *request.OrderStoreRequest) (*models.Order, error) {
	agency := models.Agency{
		Name: req.GetAgencyName(),
	}

	userName := req.GetUserName()
	note := req.GetNote()
	adminNote := req.GetAdminNote()

	order := &models.Order{
		Agency:         agency,
		RentalDate:     req.GetRentalDate(),
		GuestCount:     req.GetGuestCount(),
		TransportCount: req.GetTransportCount(),
		UserName:       &userName,
		Email:          req.GetEmail(),
		Phone:          req.GetPhone(),
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
