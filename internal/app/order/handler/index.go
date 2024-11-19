package handler

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) Index(ctx *gin.Context) {
	var req request.OrderIndexRequest

	if err := req.Validate(ctx); err != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, err)
		return
	}

	response := handler.useCase.GetOrderList(&req)

	handler.presenter.PresentOrderList(ctx, response)
}
