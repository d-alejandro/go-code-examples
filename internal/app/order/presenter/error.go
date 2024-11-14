package presenter

import (
	"net/http"

	"github.com/d-alejandro/go-code-examples/internal/app/order/resource"
	"github.com/gin-gonic/gin"
)

type ErrorPresenter struct {
	context      *gin.Context
	statusCode   int
	messageError any
}

func PresentErrorPresenter(context *gin.Context, statusCode int, messageError any) {
	errorPresenter := &ErrorPresenter{
		context:      context,
		statusCode:   statusCode,
		messageError: messageError,
	}
	errorPresenter.present()
}

func (presenter *ErrorPresenter) present() {
	var errorsString string

	switch messageError := presenter.messageError.(type) {
	case string:
		errorsString = messageError
	case error:
		errorsString = messageError.Error()
	default:
		errorsString = ""
	}

	message := http.StatusText(presenter.statusCode)
	errorResource := resource.NewErrorResource(message, errorsString)

	presenter.context.JSON(presenter.statusCode, errorResource)
}
