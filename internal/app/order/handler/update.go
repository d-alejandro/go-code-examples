package handler

import (
	"net/http"
	"strconv"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
	"github.com/gin-gonic/gin"
)

func (handler *orderHandler) Update(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, errParam := strconv.Atoi(paramID)
	if errParam != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, config.MessageInvalidID)
		return
	}

	var req request.OrderUpdateRequest

	if err := req.Validate(ctx); err != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, err)
		return
	}

	response, useCaseError := handler.useCase.Update(&req, id)
	if useCaseError != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, useCaseError)
		return
	}

	handler.presenter.PresentOrder(ctx, response)
}
