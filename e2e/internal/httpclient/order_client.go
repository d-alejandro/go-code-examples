package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/d-alejandro/go-code-examples/e2e/internal/config"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

type OrderClient interface {
	Create(*request.OrderCreationRequest) (*response.OrderResponse, *http.Response, error)
	Delete(id int) (*response.OrderResponse, *http.Response, error)
	Get(id int) (*response.OrderResponse, *http.Response, error)
	Paginate(*request.OrderPaginationRequest) (*response.OrderListResponse, *http.Response, error)
	Update(*request.OrderUpdateRequest) (*response.OrderResponse, *http.Response, error)
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

func (cli *orderClient) send(method string, reqUrl string, body io.Reader, resp any) (*http.Response, error) {
	httpRequest, err := http.NewRequest(method, reqUrl, body)

	if err != nil {
		return nil, err
	}

	if resp != nil {
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	httpResponse, respErr := cli.client.Do(httpRequest)

	if respErr != nil {
		return nil, respErr
	}

	responseBody, readErr := io.ReadAll(httpResponse.Body)

	if readErr != nil {
		return nil, readErr
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(responseBody))
	}

	err = json.Unmarshal(responseBody, resp)

	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}
