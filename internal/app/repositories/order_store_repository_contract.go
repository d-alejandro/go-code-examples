package repositories

import (
	"github.com/gin-gonic/gin"
	"time"
)

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
