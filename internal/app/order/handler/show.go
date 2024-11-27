package handler

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) Show(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, config.MessageInvalidID)
		return
	}

	order, useCaseErr := handler.useCase.GetOrder(ctx, id)
	if useCaseErr != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, useCaseErr)
		return
	}

	handler.presenter.PresentOrder(ctx, order)
}
