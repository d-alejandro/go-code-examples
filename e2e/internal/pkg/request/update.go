package request

import (
	"strconv"

	"github.com/d-alejandro/go-code-examples/e2e/internal/config"
)

type OrderUpdateRequest struct {
	ID             int
	GuestCount     int
	TransportCount int
	UserName       *string
	Email          string
	Phone          string
	Note           *string
	AdminNote      *string
}

func (request *OrderUpdateRequest) GetID() int {
	return request.ID
}

func (request *OrderUpdateRequest) GetGuestCount() string {
	return strconv.Itoa(request.GuestCount)
}

func (request *OrderUpdateRequest) GetTransportCount() string {
	return strconv.Itoa(request.TransportCount)
}

func (request *OrderUpdateRequest) GetUserName() string {
	if request.UserName != nil {
		return *request.UserName
	}
	return config.EmptyString
}

func (request *OrderUpdateRequest) GetEmail() string {
	return request.Email
}

func (request *OrderUpdateRequest) GetPhone() string {
	return request.Phone
}

func (request *OrderUpdateRequest) GetNote() string {
	if request.Note != nil {
		return *request.Note
	}
	return config.EmptyString
}

func (request *OrderUpdateRequest) GetAdminNote() string {
	if request.AdminNote != nil {
		return *request.AdminNote
	}
	return config.EmptyString
}
