package request

type OrderUpdateRequest struct {
	GuestCount     int     `form:"guest_count" binding:"required,numeric,min=1"`
	TransportCount int     `form:"transport_count" binding:"required,numeric,min=1"`
	UserName       *string `form:"user_name" binding:"omitempty,max=100"`
	Email          string  `form:"email" binding:"required,email,max=100"`
	Phone          string  `form:"phone" binding:"required,max=50"`
	Note           *string `form:"note" binding:"omitempty,max=1000"`
	AdminNote      *string `form:"admin_note" binding:"omitempty,max=1000"`
}

func (request *OrderUpdateRequest) GetGuestCount() int {
	return request.GuestCount
}

func (request *OrderUpdateRequest) GetTransportCount() int {
	return request.TransportCount
}

func (request *OrderUpdateRequest) GetUserName() *string {
	return request.UserName
}

func (request *OrderUpdateRequest) GetEmail() string {
	return request.Email
}

func (request *OrderUpdateRequest) GetPhone() string {
	return request.Phone
}

func (request *OrderUpdateRequest) GetNote() *string {
	return request.Note
}

func (request *OrderUpdateRequest) GetAdminNote() *string {
	return request.AdminNote
}
