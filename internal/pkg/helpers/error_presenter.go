package helpers

import (
	"context"
	"net/http"
)

type ErrorPresenter struct {
}

type contextExtended interface {
	context.Context
	JSON(code int, obj any)
}

func (*ErrorPresenter) PresentError(ctx interface{ contextExtended }, statusCode int, errors any) {
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
