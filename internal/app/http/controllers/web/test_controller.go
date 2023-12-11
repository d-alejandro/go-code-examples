package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
}

func NewTestController() *TestController {
	return new(TestController)
}

func (controller *TestController) Test(context *gin.Context) {
	context.String(http.StatusOK, "The server started successfully!")
}
