package middleware

import (
	"net/http"
	"product_restful_api/helper"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-KEY") == "123QWE" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := helper.NewWebResponse(http.StatusUnauthorized, "UNAUTORIZED", nil)
		helper.WriteToResponseBody(writer, webResponse)
	}
}

