package orderclient

import (
	"fmt"
	"net/http"

	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

func (client *orderClient) Get(id int) (*response.OrderResponse, error) {
	requestUrl := fmt.Sprintf("%s/api/orders/%d", client.connection.HTTPServerUrl, id)

	var orderResponse response.OrderResponse

	err := client.send(http.MethodGet, requestUrl, nil, &orderResponse)

	if err != nil {
		return nil, err
	}

	return &orderResponse, nil
}
