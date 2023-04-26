package helper

import "product_restful_api/model/web"

func NewWebResponse(code int, status string, data interface{}) web.WebResponse {
	return web.WebResponse{
		Code: code,
		Status: status,
		Data: data,
	}
}