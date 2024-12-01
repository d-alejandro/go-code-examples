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

	id, err := strconv.Atoi(paramID)
	if err != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, config.MessageInvalidID)
		return
	}

	var req request.OrderUpdateRequest

	err = handler.validator.ValidateForm(ctx, &req)
	if err != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, err)
		return
	}

	order, useCaseErr := handler.useCase.Update(ctx, &req, id)
	if useCaseErr != nil {
		handler.presenter.PresentError(ctx, http.StatusBadRequest, useCaseErr)
		return
	}

	handler.presenter.PresentOrder(ctx, order)
}
