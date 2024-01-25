package use_cases

import "github.com/d-alejandro/go-code-examples/internal/app/models"

type OrderIndexRepositoryInterface interface {
	Make(interface{ PaginationDTOInterface }) []models.Order
}

type PaginationDTOInterface interface {
	GetSortColumn() string
	GetSortType() string
	GetLimitValue() int
	GetOffsetValue() int
}
