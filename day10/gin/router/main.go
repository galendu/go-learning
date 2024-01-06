package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func boy(c *gin.Context) {

	c.String(http.StatusOK, "hi boy")
}

func girl(c *gin.Context) {

	c.String(http.StatusOK, "hi girl")
}
func main() {
	engine := gin.Default()
	engine.GET("/", boy)
	engine.POST("/", girl)

	//路由分组
	oldVersion := engine.Group("/v1")
	oldVersion.GET("/student", boy)
	oldVersion.GET("/teacher", boy)

	newVersion := engine.Group("v2")
	newVersion.GET("/student", girl)
	newVersion.GET("/teacher", girl)
	engine.Run(":5656")
}
