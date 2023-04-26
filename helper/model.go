package helper

import (
	"product_restful_api/model/domain"
	"product_restful_api/model/web"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ToProductResponse(product *domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse

	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(&product))
	}

	return productResponses
}