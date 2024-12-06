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

func (client *orderClient) send(method string, reqURL string, body io.Reader, resp any) (err error) {
	var httpRequest *http.Request

	if httpRequest, err = http.NewRequest(method, reqURL, body); err != nil {
		return
	}

	if resp != nil {
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	var httpResponse *http.Response

	if httpResponse, err = client.client.Do(httpRequest); err != nil {
		return
	}

	defer func() {
		if bodyCloseErr := httpResponse.Body.Close(); bodyCloseErr != nil {
			if err != nil {
				return
			}
			err = bodyCloseErr
		}
	}()

	var responseBody []byte

	if responseBody, err = io.ReadAll(httpResponse.Body); err != nil {
		return
	}

	if httpResponse.StatusCode != http.StatusOK {
		err = helpers.NewHTTPError(httpResponse.StatusCode, string(responseBody))
		return
	}

	err = json.Unmarshal(responseBody, resp)
	return
}
