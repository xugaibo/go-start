package response

import (
	"go-start/core/bizcode"
)

type ResultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(data interface{}) ResultResponse {
	response := ResultResponse{}
	response.Code = bizcode.Success.Code()
	response.Data = data
	return response
}

func NotOk(code int, message string) ResultResponse {
	response := ResultResponse{}
	response.Code = code
	response.Message = message
	return response
}
