package exception

import (
	"net/http"
	"product_restful_api/helper"

	"github.com/go-playground/validator/v10"
)

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := helper.NewWebResponse(http.StatusBadRequest, "BAD REQUEST", exception.Error())
		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := helper.NewWebResponse(http.StatusNotFound, "NOT FOUND", exception.Error)
		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func InternalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := helper.NewWebResponse(http.StatusInternalServerError, "INTERNAL SERVER ERROR", err)
	helper.WriteToResponseBody(writer, webResponse)
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	InternalServerError(writer, request, err)
}