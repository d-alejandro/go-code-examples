package presenters

import (
	"github.com/d-alejandro/go-code-examples/internal/app/http/resources"
	"github.com/gin-gonic/gin"
	"net/http"
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

	errorResource := resources.ErrorResource{
		Message: http.StatusText(presenter.statusCode),
		Errors:  messageError,
	}

	presenter.context.JSON(presenter.statusCode, errorResource)
}
