package api

import (
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/gin-gonic/gin"
)

type OrderIndexUseCaseInterface interface {
	Execute(interface{ OrderIndexRequestInterface }) []*models.Order
}

type OrderIndexRequestInterface interface {
	Validate(*gin.Context) error
}

type OrderIndexPresenterInterface interface {
	Present(*gin.Context, []*models.Order)
}
