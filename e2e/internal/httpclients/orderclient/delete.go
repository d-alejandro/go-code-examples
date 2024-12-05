package orderclient

import (
	"fmt"
	"net/http"

	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

func (cli *orderClient) Delete(id int) (*response.OrderResponse, *http.Response, error) {
	requestUrl := fmt.Sprintf("%s/api/orders/%d", cli.connection.HTTPServerUrl, id)

	var orderResponse response.OrderResponse

	httpResponse, err := cli.send(http.MethodDelete, requestUrl, nil, &orderResponse)

	return &orderResponse, httpResponse, err
}
