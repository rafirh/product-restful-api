package service

import (
	"context"
	"product_restful_api/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductRequestCreate) web.ProductResponse
	Update(ctx context.Context, request web.ProductRequestUpdate) web.ProductResponse
	Delete(ctx context.Context, productId int)
	FindById(ctx context.Context, productId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}