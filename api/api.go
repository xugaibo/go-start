package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-start/bizcode"
	"go-start/bizerror"
	"go-start/models/response"
	"net/http"
	"runtime"
)

type Api struct {
	Context *gin.Context
}

func (a *Api) Ok() {
	a.Context.AbortWithStatusJSON(http.StatusOK, response.Ok(true))
}

func (a *Api) Success(data any) {
	a.Context.AbortWithStatusJSON(http.StatusOK, response.Ok(data))
}

func (a *Api) ClientError(err error) {
	fmt.Println("client error:")
	fmt.Println(err)
	fmt.Println(getCurrentGoroutineStack())
	withError := response.NotOk(bizcode.ClientError.Code(), fmt.Sprint(err))
	a.Context.AbortWithStatusJSON(http.StatusOK, withError)
}

func (a *Api) Error(bizError bizerror.BizError) {
	withError := response.NotOk(bizError.Code, bizError.Error())
	a.Context.AbortWithStatusJSON(http.StatusOK, withError)
}

func (a *Api) MakeContext(c *gin.Context) *Api {
	a.Context = c
	return a
}

func getCurrentGoroutineStack() string {
	var buf [1000]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}

func (a *Api) ErrorHandler() {
	if err := recover(); err != nil {
		if _, ok := err.(bizerror.BizError); ok {
			bizError, jsonErr := bizerror.Parse(err)
			if jsonErr != nil {
				a.Error(bizerror.Wrap(jsonErr))
			}

			fmt.Println("biz error:")
			fmt.Println(err)

			a.Error(*bizError)
		} else if _, ok := err.(error); ok {
			fmt.Println("server error:")
			fmt.Println(err)
			fmt.Println(getCurrentGoroutineStack())

			a.Error(bizerror.Biz(bizcode.ServerError))
		}
	}
}
