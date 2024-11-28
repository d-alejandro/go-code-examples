package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ValidationHelper struct {
}

func NewValidationHelper() *ValidationHelper {
	return &ValidationHelper{}
}

func (validator *ValidationHelper) ValidateForm(ctx *gin.Context, request any) error {
	return ctx.ShouldBindWith(request, binding.Form)
}

func (validator *ValidationHelper) ValidateQuery(ctx *gin.Context, request any) error {
	return ctx.ShouldBindWith(request, binding.Query)
}
