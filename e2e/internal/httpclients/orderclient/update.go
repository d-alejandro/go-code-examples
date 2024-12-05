package orderclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

func (client *orderClient) Update(req *request.OrderUpdateRequest) (*response.OrderResponse, error) {
	formData := url.Values{}

	formData.Set("guest_count", req.GetGuestCount())
	formData.Set("transport_count", req.GetTransportCount())
	formData.Set("user_name", req.GetUserName())
	formData.Set("email", req.GetEmail())
	formData.Set("phone", req.GetPhone())
	formData.Set("note", req.GetNote())
	formData.Set("admin_note", req.GetAdminNote())

	encodedFormData := formData.Encode()
	requestBody := strings.NewReader(encodedFormData)

	requestUrl := fmt.Sprintf("%s/api/orders/%d", client.connection.HTTPServerUrl, req.ID)

	var orderResponse response.OrderResponse

	err := client.send(http.MethodPut, requestUrl, requestBody, &orderResponse)

	if err != nil {
		return nil, err
	}

	return &orderResponse, nil
}
