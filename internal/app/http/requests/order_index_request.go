package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderIndexRequest struct {
	Start      int    `form:"start" binding:"required|eq=0,numeric,min=0"`
	End        int    `form:"end" binding:"required,numeric,min=1"`
	SortColumn string `form:"sort_column" binding:"required,oneof=id status rental_date guest_count created_at"`
	SortType   string `form:"sort_type" binding:"required,oneof=asc desc"`
}

func (request *OrderIndexRequest) Validate(context *gin.Context) error {
	return context.ShouldBindWith(request, binding.Query)
}

func (request *OrderIndexRequest) GetStart() int {
	return request.Start
}

func (request *OrderIndexRequest) GetEnd() int {
	return request.End
}

func (request *OrderIndexRequest) GetSortColumn() string {
	return request.SortColumn
}

func (request *OrderIndexRequest) GetSortType() string {
	return request.SortType
}
