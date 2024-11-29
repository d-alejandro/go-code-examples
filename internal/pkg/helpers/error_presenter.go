package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorPresenter struct {
	rendering RenderingHelper
}

func (presenter *ErrorPresenter) PresentError(ctx *gin.Context, statusCode int, errors any) {
	statusCodeText := http.StatusText(statusCode)

	responseBody := &struct {
		Message string `json:"message"`
		Errors  string `json:"errors"`
	}{
		Message: statusCodeText,
		Errors:  fmt.Sprintf("%v", errors),
	}

	presenter.rendering.RenderJSON(ctx, statusCode, responseBody)
}
