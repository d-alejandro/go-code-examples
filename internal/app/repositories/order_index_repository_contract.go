package repositories

type PaginationDTOInterface interface {
	GetSortColumn() string
	GetSortType() string
	GetLimitValue() int
	GetOffsetValue() int
}
