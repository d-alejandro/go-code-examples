package handler

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) Index(ctx *gin.Context) {
	var req request.OrderIndexRequest

	if err := handler.validator.ValidateQuery(ctx, &req); err != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, err)
		return
	}

	orders := handler.useCase.GetOrderList(ctx, &req)

	handler.presenter.PresentOrderList(ctx, orders)
}
