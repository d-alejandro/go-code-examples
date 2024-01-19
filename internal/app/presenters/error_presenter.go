package presenters

import (
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
	var messageError any

	switch (presenter.messageError).(type) {
	case string:
		messageError = presenter.messageError
	default:
		messageError = (presenter.messageError).(error).Error()
	}

	object := gin.H{
		"error": messageError,
	}

	presenter.context.JSON(presenter.statusCode, object)
}
