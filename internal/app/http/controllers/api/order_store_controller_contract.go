package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/gin-gonic/gin"
	"time"
)

type OrderStoreUseCaseInterface interface {
	Execute(interface{ OrderStoreRequestInterface }) (*models.Order, error)
}

type OrderStoreRequestInterface interface {
	Validate(*gin.Context) error
	GetAgencyName() string
	GetRentalDate() *time.Time
	GetGuestCount() int
	GetTransportCount() int
	GetUserName() string
	GetEmail() string
	GetPhone() string
	GetNote() string
	GetAdminNote() string
}

type OrderStorePresenterInterface interface {
	Present(*gin.Context, *models.Order)
}
