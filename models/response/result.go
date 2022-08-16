package response

import (
	"go-start/core/bizcode"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(data interface{}) Result {
	response := Result{}
	response.Code = bizcode.Success.Code()
	response.Data = data
	return response
}

func NotOk(code int, message string) Result {
	response := Result{}
	response.Code = code
	response.Message = message
	return response
}

func Biz(biz bizcode.BizCode) Result {
	response := Result{}
	response.Code = biz.Code()
	response.Message = biz.String()
	return response
}
