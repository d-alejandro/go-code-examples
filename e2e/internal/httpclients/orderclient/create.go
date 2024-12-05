package orderclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

func (cli *orderClient) Create(req *request.OrderCreationRequest) (*response.OrderResponse, *http.Response, error) {
	formData := url.Values{}

	formData.Set("agency_name", req.GetAgencyName())
	formData.Set("rental_date", req.GetRentalDate())
	formData.Set("guest_count", req.GetGuestCount())
	formData.Set("transport_count", req.GetTransportCount())
	formData.Set("user_name", req.GetUserName())
	formData.Set("email", req.GetEmail())
	formData.Set("phone", req.GetPhone())
	formData.Set("note", req.GetNote())
	formData.Set("admin_note", req.GetAdminNote())

	encodedFormData := formData.Encode()
	requestBody := strings.NewReader(encodedFormData)

	requestUrl := fmt.Sprintf("%s/api/orders", cli.connection.HTTPServerUrl)

	var orderResponse response.OrderResponse

	httpResponse, err := cli.send(http.MethodPost, requestUrl, requestBody, &orderResponse)

	return &orderResponse, httpResponse, err
}
