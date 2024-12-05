package request

type OrderPaginationRequest struct {
	Start      int
	End        int
	SortColumn string
	SortType   string
	IDs        []int
}

func (request *OrderPaginationRequest) GetStart() int {
	return request.Start
}

func (request *OrderPaginationRequest) GetEnd() int {
	return request.End
}

func (request *OrderPaginationRequest) GetSortColumn() string {
	return request.SortColumn
}

func (request *OrderPaginationRequest) GetSortType() string {
	return request.SortType
}

func (request *OrderPaginationRequest) GetIDs() []int {
	return request.IDs
}
