package response

type OrderListResponse struct {
	Data []struct {
		ID             int     `json:"id"`
		AgencyName     string  `json:"agency_name"`
		Status         string  `json:"status"`
		IsConfirmed    bool    `json:"is_confirmed"`
		IsChecked      bool    `json:"is_checked"`
		RentalDate     string  `json:"rental_date"`
		UserName       *string `json:"user_name"`
		TransportCount int     `json:"transport_count"`
		GuestCount     int     `json:"guest_count"`
		AdminNote      *string `json:"admin_note"`
	} `json:"data"`
}
