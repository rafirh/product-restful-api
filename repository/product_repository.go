package repository

import (
	"context"
	"database/sql"
	"product_restful_api/model/domain"
)

type ProductRepository interface {
	FindAll(ctx context.Context, db *sql.DB, product domain.Product) []domain.Product
	FindById(ctx context.Context, db *sql.DB, productId int) domain.Product
	Create(ctx context.Context, db *sql.DB, product domain.Product) domain.Product
	Update(ctx context.Context, db *sql.DB, product domain.Product) domain.Product
	Delete(ctx context.Context, db *sql.DB, productId int) 
}