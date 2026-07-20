package route

import (
	"shop/api/handler"

	"github.com/gin-gonic/gin"
)

func Generate(r *gin.RouterGroup) {
	r.POST("/", handler.HandleCreateCategory)
	r.GET("/", handler.HandleGetCategory)
	r.GET("/:id", nil)
	r.PUT("/", nil)
	r.PATCH("/:id/status", nil)
	r.DELETE("/:id", nil)

}
