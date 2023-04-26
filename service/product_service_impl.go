package service

import (
	"context"
	"database/sql"
	"product_restful_api/exception"
	"product_restful_api/helper"
	"product_restful_api/model/domain"
	"product_restful_api/model/web"
	"product_restful_api/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service ProductServiceImpl) Create(ctx context.Context, request web.ProductRequestCreate) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name: request.Name,
		Price: request.Price,
	}

	product = service.ProductRepository.Save(ctx, tx, product)
	return helper.ToProductResponse(&product)
}

func (service ProductServiceImpl) Update(ctx context.Context, request web.ProductRequestUpdate) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = request.Name
	product.Price = request.Price

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(&product)
}

func (service ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(&product)
}

func (service ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)
	productResponses := helper.ToProductResponses(products)

	return productResponses
}
