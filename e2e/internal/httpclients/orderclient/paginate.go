package orderclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

func (client *orderClient) Paginate(req *request.OrderPaginationRequest) (*response.OrderListResponse, error) {
	urlData := url.Values{}

	urlData.Set("start", req.GetStart())
	urlData.Set("end", req.GetEnd())
	urlData.Set("sort_column", req.GetSortColumn())
	urlData.Set("sort_type", req.GetSortType())

	encodedUrlData := urlData.Encode()

	requestUrl := fmt.Sprintf("%s/api/orders?%s%s", client.connection.HTTPServerUrl, encodedUrlData, req.GetIDs())

	var orderResponse response.OrderListResponse

	err := client.send(http.MethodGet, requestUrl, nil, &orderResponse)

	if err != nil {
		return nil, err
	}

	return &orderResponse, nil
}
