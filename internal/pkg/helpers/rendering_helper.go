package helpers

import "github.com/gin-gonic/gin"

type RenderingHelper interface {
	RenderJSON(ctx *gin.Context, code int, obj any)
}

type renderingHelper struct {
}

func NewRenderingHelper() RenderingHelper {
	return &renderingHelper{}
}

func (*renderingHelper) RenderJSON(ctx *gin.Context, code int, obj any) {
	ctx.JSON(code, obj)
}
