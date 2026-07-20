package handler

import (
	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	
}

func HandleCreateCategory(c *gin.Context) {
	var req CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"name":   req.Name,
	})
}

func HandleGetCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
}

func HandleEditCategory(c *gin.Context) {

}

func HandleEditStatusCategory(c *gin.Context) {

}

func HandleDeleteCategory(c *gin.Context) {
}
