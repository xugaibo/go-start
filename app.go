package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-start/apis"
	"go-start/db"
	"go-start/middleware"
)

func main() {
	db.InitDb()

	r := gin.Default()
	api := apis.Article{}

	user := apis.User{}
	r.POST("/user", user.Create)

	token := apis.Token{}
	r.POST("/token", token.Create)

	r2 := r.Group("/")
	r2.Use(middleware.CheckLogin())
	// apis
	r2.GET("/article", api.List)
	r2.POST("/article", api.Create)
	r2.DELETE("/article/:id", api.Delete)
	r2.PUT("/article", api.Update)

	err := r.Run()
	if err != nil {
		fmt.Println("server start fail")
		panic(err)
	}
}
