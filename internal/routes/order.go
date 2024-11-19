package routes

import "github.com/gin-gonic/gin"

const (
	rootPath    = "/"
	idParamPath = "/:id"
)

type OrderHandler interface {
	Index(*gin.Context)
	Show(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Destroy(*gin.Context)
}

func RegisterOrderHandles(engine *gin.Engine, handler OrderHandler) {
	api := engine.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.GET(rootPath, handler.Index)
			orders.GET(idParamPath, handler.Show)
			orders.POST(rootPath, handler.Store)
			orders.PUT(idParamPath, handler.Update)
			orders.DELETE(idParamPath, handler.Destroy)
		}
	}
}
