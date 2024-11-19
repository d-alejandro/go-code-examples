package handler

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) Store(ctx *gin.Context) {
	var req request.OrderStoreRequest

	if err := req.Validate(ctx); err != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, err)
		return
	}

	response, useCaseError := handler.useCase.Store(&req)
	if useCaseError != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, useCaseError)
		return
	}

	handler.presenter.PresentOrder(ctx, response)
}
