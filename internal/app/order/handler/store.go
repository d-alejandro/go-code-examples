package handler

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) Store(ctx *gin.Context) {
	var req request.OrderStoreRequest

	if err := handler.validator.ValidateForm(ctx, &req); err != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, err)
		return
	}

	order, useCaseErr := handler.useCase.Create(ctx, &req)
	if useCaseErr != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, useCaseErr)
		return
	}

	handler.presenter.PresentOrder(ctx, order)
}
