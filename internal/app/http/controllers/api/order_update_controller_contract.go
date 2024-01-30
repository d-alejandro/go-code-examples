package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/gin-gonic/gin"
)

type OrderUpdateUseCaseInterface interface {
	Execute(orderUpdateRequest interface{ OrderUpdateRequestInterface }, id int) (*models.Order, error)
}

type OrderUpdateRequestInterface interface {
	Validate(*gin.Context) error
	GetGuestCount() int
	GetTransportCount() int
	GetUserName() string
	GetEmail() string
	GetPhone() string
	GetNote() string
	GetAdminNote() string
}

type OrderUpdatePresenterInterface interface {
	Present(*gin.Context, *models.Order)
}
