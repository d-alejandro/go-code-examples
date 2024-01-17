package requests

type OrderIndexRequest struct {
	Start      int    `form:"start" binding:"required|eq=0,numeric,min=0"`
	End        int    `form:"end" binding:"required,numeric,min=1"`
	SortColumn string `form:"sort_column" binding:"required"`
	SortType   string `form:"sort_type" binding:"required"`
}
