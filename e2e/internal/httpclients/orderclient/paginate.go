package orderclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/request"
	"github.com/d-alejandro/go-code-examples/e2e/internal/pkg/response"
)

func (cli *orderClient) Paginate(req *request.OrderPaginationRequest) (
	*response.OrderListResponse,
	*http.Response,
	error,
) {
	urlData := url.Values{}

	urlData.Set("start", req.GetStart())
	urlData.Set("end", req.GetEnd())
	urlData.Set("sort_column", req.GetSortColumn())
	urlData.Set("sort_type", req.GetSortType())

	encodedUrlData := urlData.Encode()

	requestUrl := fmt.Sprintf("%s/api/orders?%s%s", cli.connection.HTTPServerUrl, encodedUrlData, req.GetIDs())

	var orderResponse response.OrderListResponse

	httpResponse, err := cli.send(http.MethodGet, requestUrl, nil, &orderResponse)

	return &orderResponse, httpResponse, err
}
