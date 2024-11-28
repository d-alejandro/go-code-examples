package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ValidationHelper interface {
	ValidateForm(ctx *gin.Context, request any) error
	ValidateQuery(ctx *gin.Context, request any) error
}

type validationHelper struct {
}

func NewValidationHelper() ValidationHelper {
	return &validationHelper{}
}

func (validator *validationHelper) ValidateForm(ctx *gin.Context, request any) error {
	return ctx.ShouldBindWith(request, binding.Form)
}

func (validator *validationHelper) ValidateQuery(ctx *gin.Context, request any) error {
	return ctx.ShouldBindWith(request, binding.Query)
}
