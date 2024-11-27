package request

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderStoreRequest struct {
	AgencyName     string    `form:"agency_name" binding:"required,max=200"`
	RentalDate     time.Time `form:"rental_date" binding:"required,gte" time_format:"02-01-2006"`
	GuestCount     int       `form:"guest_count" binding:"required,numeric,min=1"`
	TransportCount int       `form:"transport_count" binding:"required,numeric,min=1"`
	UserName       *string   `form:"user_name" binding:"omitempty,max=100"`
	Email          string    `form:"email" binding:"required,email,max=100"`
	Phone          string    `form:"phone" binding:"required,max=50"`
	Note           *string   `form:"note" binding:"omitempty,max=1000"`
	AdminNote      *string   `form:"admin_note" binding:"omitempty,max=1000"`
}

func (request *OrderStoreRequest) Validate(context *gin.Context) error {
	return context.ShouldBindWith(request, binding.Form)
}

func (request *OrderStoreRequest) GetAgencyName() string {
	return request.AgencyName
}

func (request *OrderStoreRequest) GetRentalDate() time.Time {
	return request.RentalDate
}

func (request *OrderStoreRequest) GetGuestCount() int {
	return request.GuestCount
}

func (request *OrderStoreRequest) GetTransportCount() int {
	return request.TransportCount
}

func (request *OrderStoreRequest) GetUserName() *string {
	return request.UserName
}

func (request *OrderStoreRequest) GetEmail() string {
	return request.Email
}

func (request *OrderStoreRequest) GetPhone() string {
	return request.Phone
}

func (request *OrderStoreRequest) GetNote() *string {
	return request.Note
}

func (request *OrderStoreRequest) GetAdminNote() *string {
	return request.AdminNote
}
