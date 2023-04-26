package controller

import (
	"net/http"
	"product_restful_api/helper"
	"product_restful_api/model/web"
	"product_restful_api/service"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	service service.ProductService
}

func (controller ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var productRequest web.ProductRequestCreate
	helper.ReadFromRequestBody(request, &productRequest)

	productResponse := controller.service.Create(request.Context(), productRequest)
	webResponse := helper.NewWebResponse(200, "OK", productResponse)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var productRequest web.ProductRequestUpdate
	helper.ReadFromRequestBody(request, productRequest)

	productRequest.Id = helper.GetParamInt(&params, "id")

	productResponse := controller.service.Update(request.Context(), productRequest)
	webResponse := helper.NewWebResponse(200, "OK", productResponse)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := helper.GetParamInt(&params, "id")

	controller.service.Delete(request.Context(), productId)
	webResponse := helper.NewWebResponse(200, "OK", nil)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := helper.GetParamInt(&params, "id")
	
	productResponse := controller.service.FindById(request.Context(), productId)
	webResponse := helper.NewWebResponse(200, "OK", productResponse)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponses := controller.service.FindAll(request.Context())
	webResponse := helper.NewWebResponse(200, "OK", productResponses)

	helper.WriteToResponseBody(writer, webResponse)
}