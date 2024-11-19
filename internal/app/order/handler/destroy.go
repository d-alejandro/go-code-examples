package handler

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) Destroy(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, errParam := strconv.Atoi(paramID)
	if errParam != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, config.MessageInvalidID)
		return
	}

	response, useCaseError := handler.useCase.Delete(id)
	if useCaseError != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, useCaseError)
		return
	}

	handler.presenter.PresentOrder(ctx, response)
}
