package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/gin-gonic/gin"
)

type OrderIndexUseCaseInterface interface {
	Execute(interface{ OrderIndexRequestInterface }) []models.Order
}

type OrderIndexRequestInterface interface {
	GetStart() int
	GetEnd() int
	GetSortColumn() string
	GetSortType() string
}

type OrderIndexPresenterInterface interface {
	Present(*gin.Context, []models.Order)
}
