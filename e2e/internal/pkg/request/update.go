package request

import "strconv"

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
	return ""
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
	return ""
}

func (request *OrderUpdateRequest) GetAdminNote() string {
	if request.AdminNote != nil {
		return *request.AdminNote
	}
	return ""
}
