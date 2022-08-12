package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-start/apis"
	"go-start/core/context"
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

	err := r.Run()
	if err != nil {
		fmt.Println("server start fail")
		panic(err)
	}
}
