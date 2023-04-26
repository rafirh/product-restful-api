package main

import (
	"net/http"
	"product_restful_api/app"
	"product_restful_api/controller"
	"product_restful_api/helper"
	"product_restful_api/middleware"
	"product_restful_api/repository"
	"product_restful_api/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := app.NewRouter(productController)

	server := http.Server{
		Addr: ":3333",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}