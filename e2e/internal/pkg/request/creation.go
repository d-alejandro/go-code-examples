package request

import "time"

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

func (request *OrderCreationRequest) GetRentalDate() time.Time {
	return request.RentalDate
}

func (request *OrderCreationRequest) GetGuestCount() int {
	return request.GuestCount
}

func (request *OrderCreationRequest) GetTransportCount() int {
	return request.TransportCount
}

func (request *OrderCreationRequest) GetUserName() *string {
	return request.UserName
}

func (request *OrderCreationRequest) GetEmail() string {
	return request.Email
}

func (request *OrderCreationRequest) GetPhone() string {
	return request.Phone
}

func (request *OrderCreationRequest) GetNote() *string {
	return request.Note
}

func (request *OrderCreationRequest) GetAdminNote() *string {
	return request.AdminNote
}
