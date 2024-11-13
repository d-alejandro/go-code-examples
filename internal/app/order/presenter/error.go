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
	var messageError string

	switch presenter.messageError.(type) {
	case string:
		messageError = presenter.messageError.(string)
	case error:
		messageError = presenter.messageError.(error).Error()
	default:
		messageError = ""
	}

	message := http.StatusText(presenter.statusCode)
	errorResource := resource.NewErrorResource(message, messageError)

	presenter.context.JSON(presenter.statusCode, errorResource)
}
