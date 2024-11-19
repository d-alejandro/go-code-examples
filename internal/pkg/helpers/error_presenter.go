package helpers

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/pkg/resource"
	"github.com/gin-gonic/gin"
)

type ErrorPresenter struct {
}

func (*ErrorPresenter) PresentError(ctx *gin.Context, statusCode int, errors any) {
	var errorsText string

	switch err := errors.(type) {
	case string:
		errorsText = err
	case error:
		errorsText = err.Error()
	default:
		errorsText = "unknown error"
	}

	statusCodeText := http.StatusText(statusCode)
	errorResource := resource.NewErrorResource(statusCodeText, errorsText)

	ctx.JSON(statusCode, errorResource)
}
