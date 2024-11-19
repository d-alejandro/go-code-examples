package presenter

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

type OrderPresenter interface {
	PresentOrder(*gin.Context, *models.Order)
	PresentOrderList(*gin.Context, []*models.Order)
	errorPresenter
}

type errorPresenter interface {
	PresentError(ctx *gin.Context, statusCode int, errors any)
}

type orderPresenter struct {
	helpers.ErrorPresenter
}

func NewOrderPresenter() OrderPresenter {
	return &orderPresenter{}
}
