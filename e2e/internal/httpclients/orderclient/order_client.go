package orderclient

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/d-alejandro/go-code-examples/e2e/internal/config"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

type OrderClient interface {
	Create(*request.OrderCreationRequest) (*response.OrderResponse, error)
	Delete(id int) (*response.OrderResponse, error)
	Get(id int) (*response.OrderResponse, error)
	Paginate(*request.OrderPaginationRequest) (*response.OrderListResponse, error)
	Update(*request.OrderUpdateRequest) (*response.OrderResponse, error)
}

type orderClient struct {
	client     *http.Client
	connection *config.Connection
}

func NewOrderClient() OrderClient {
	connection := config.NewConnection()

	client := &http.Client{
		Timeout: connection.HTTPClientTimeout,
	}

	return &orderClient{
		client:     client,
		connection: connection,
	}
}

func (client *orderClient) send(method string, reqUrl string, body io.Reader, resp any) error {
	httpRequest, err := http.NewRequest(method, reqUrl, body)

	if err != nil {
		return err
	}

	if resp != nil {
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	httpResponse, respErr := client.client.Do(httpRequest)

	if respErr != nil {
		return respErr
	}

	responseBody, readErr := io.ReadAll(httpResponse.Body)

	if readErr != nil {
		return readErr
	}

	if httpResponse.StatusCode != http.StatusOK {
		return helpers.NewHTTPError(httpResponse.StatusCode, string(responseBody))
	}

	return json.Unmarshal(responseBody, resp)
}
