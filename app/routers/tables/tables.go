package tables

import (
	"github.com/gin-gonic/gin"
)

func TablesRegisterRoutes(router *gin.RouterGroup) {
	router.GET("/:tablename",TableTest)
}

func TableTest(c *gin.Context) {
	tablename := c.Param("tablename")
	c.JSON(200, gin.H{
		"Tablename" : tablename,
	})
}