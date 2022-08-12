package main

import (
	"github.com/gin-gonic/gin"
	"go-start/apis"
	"go-start/context"
)

func main() {
	context.InitDb()

	r := gin.Default()
	api := apis.Article{}

	// apis
	r.GET("/article", api.List)
	r.POST("/article", api.Create)
	r.DELETE("/article/:id", api.Delete)
	r.PUT("/article", api.Update)

	r.Run()
}
