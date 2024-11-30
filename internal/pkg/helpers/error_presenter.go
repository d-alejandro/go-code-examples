package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorPresenter struct {
	rendering RenderingHelper
}

type ErrorResponse struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

func (presenter *ErrorPresenter) SetRenderingHelper(rendering RenderingHelper) {
	presenter.rendering = rendering
}

func (presenter *ErrorPresenter) PresentError(ctx *gin.Context, statusCode int, errors any) {
	statusCodeText := http.StatusText(statusCode)

	errorResponse := &ErrorResponse{
		Message: statusCodeText,
		Errors:  fmt.Sprintf("%v", errors),
	}

	presenter.rendering.RenderJSON(ctx, statusCode, errorResponse)
}
