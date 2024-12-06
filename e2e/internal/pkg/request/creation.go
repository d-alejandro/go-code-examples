package request

import (
	"strconv"
	"time"

	"github.com/d-alejandro/go-code-examples/e2e/internal/config"
)

type OrderCreationRequest struct {
	AgencyName     string
	RentalDate     time.Time
	GuestCount     int
	TransportCount int
	UserName       *string
	Email          string
	Phone          string
	Note           *string
	AdminNote      *string
}

func (request *OrderCreationRequest) GetAgencyName() string {
	return request.AgencyName
}

func (request *OrderCreationRequest) GetRentalDate() string {
	return request.RentalDate.Format(config.DateLayout)
}

func (request *OrderCreationRequest) GetGuestCount() string {
	return strconv.Itoa(request.GuestCount)
}

func (request *OrderCreationRequest) GetTransportCount() string {
	return strconv.Itoa(request.TransportCount)
}

func (request *OrderCreationRequest) GetUserName() string {
	if request.UserName != nil {
		return *request.UserName
	}
	return config.EmptyString
}

func (request *OrderCreationRequest) GetEmail() string {
	return request.Email
}

func (request *OrderCreationRequest) GetPhone() string {
	return request.Phone
}

func (request *OrderCreationRequest) GetNote() string {
	if request.Note != nil {
		return *request.Note
	}
	return config.EmptyString
}

func (request *OrderCreationRequest) GetAdminNote() string {
	if request.AdminNote != nil {
		return *request.AdminNote
	}
	return config.EmptyString
}
