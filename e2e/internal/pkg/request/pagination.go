package request

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/d-alejandro/go-code-examples/e2e/internal/config"
)

type OrderPaginationRequest struct {
	Start      int
	End        int
	SortColumn string
	SortType   string
	IDs        []int
}

func (request *OrderPaginationRequest) GetStart() string {
	return strconv.Itoa(request.Start)
}

func (request *OrderPaginationRequest) GetEnd() string {
	return strconv.Itoa(request.End)
}

func (request *OrderPaginationRequest) GetSortColumn() string {
	return request.SortColumn
}

func (request *OrderPaginationRequest) GetSortType() string {
	return request.SortType
}

func (request *OrderPaginationRequest) GetIDs() string {
	if request.IDs == nil {
		return config.EmptyString
	}

	var urlData strings.Builder

	for _, value := range request.IDs {
		urlData.WriteString(fmt.Sprintf("&id[]=%d", value))
	}

	return urlData.String()
}
