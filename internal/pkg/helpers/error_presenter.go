package helpers

import (
	"net/http"

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

	responseBody := &struct {
		Message string `json:"message"`
		Errors  string `json:"errors"`
	}{
		Message: statusCodeText,
		Errors:  errorsText,
	}

	ctx.JSON(statusCode, responseBody)
}
