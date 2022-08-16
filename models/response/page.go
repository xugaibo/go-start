package response

import "go-start/models/request"

type Page struct {
	Total      int64       `json:"total"`
	Data       interface{} `json:"data"`
	PageSize   int         `json:"pageSize"`
	PageNumber int         `json:"pageNumber"`
	Pages      int         `json:"pages"`
}

func NewPage(data interface{}, total int64, request request.Page) Page {
	response := Page{}
	response.Total = total
	response.Data = data
	response.PageSize = request.PageSize
	response.PageNumber = request.PageNumber
	response.Pages = getPages(request.PageSize, total)
	return response
}

func getPages(pageSize int, total int64) int {
	page := int64(pageSize)
	result := int(total / page)
	r := total % page
	if r == 0 {
		return result
	} else {
		return result + 1
	}
}
