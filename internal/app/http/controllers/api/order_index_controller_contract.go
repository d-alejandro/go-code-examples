package api

import "github.com/d-alejandro/go-code-examples/internal/app/models"

type OrderIndexUseCaseInterface interface {
	Execute(interface{ OrderIndexRequestInterface }) []models.Order
}

type OrderIndexRequestInterface interface {
	GetStart() int
	GetEnd() int
	GetSortColumn() string
	GetSortType() string
}
