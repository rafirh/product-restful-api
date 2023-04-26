package app

import (
	"product_restful_api/controller"
	"product_restful_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()
	
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:id", productController.Update)
	router.DELETE("/api/products/:id", productController.Delete)
	router.GET("/api/products/:id", productController.FindById)
	router.GET("/api/products", productController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}