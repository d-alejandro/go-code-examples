package helpers

import (
	"context"
	"fmt"
	"net/http"
)

type ErrorPresenter struct {
}

type contextExtended interface {
	context.Context
	JSON(code int, obj any)
}

func (*ErrorPresenter) PresentError(ctx interface{ contextExtended }, statusCode int, errors any) {
	statusCodeText := http.StatusText(statusCode)

	responseBody := &struct {
		Message string `json:"message"`
		Errors  string `json:"errors"`
	}{
		Message: statusCodeText,
		Errors:  fmt.Sprintf("%v", errors),
	}

	ctx.JSON(statusCode, responseBody)
}
