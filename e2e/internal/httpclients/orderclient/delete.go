package orderclient

import (
	"fmt"
	"net/http"

	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

func (client *orderClient) Delete(id int) (*response.OrderResponse, error) {
	requestURL := fmt.Sprintf("%s/api/orders/%d", client.connection.HTTPServerURL, id)

	var orderResponse response.OrderResponse

	err := client.send(http.MethodDelete, requestURL, nil, &orderResponse)

	if err != nil {
		return nil, err
	}

	return &orderResponse, nil
}
