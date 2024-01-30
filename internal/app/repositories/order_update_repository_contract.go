package repositories

import "github.com/gin-gonic/gin"

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
