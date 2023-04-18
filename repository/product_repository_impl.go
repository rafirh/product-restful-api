package repository

import (
	"context"
	"database/sql"
	"product_restful_api/helper"
	"product_restful_api/model/domain"
	"time"
)

type ProductRepositoryImpl struct {}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) FindAll(ctx context.Context, db *sql.DB, product domain.Product) []domain.Product {
	SQL := "SELECT id, name, price, created_at, updated_at FROM product"
	rows, err := db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product 
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return products
}

func (repository ProductRepositoryImpl) FindById(ctx context.Context, db *sql.DB, productId int) domain.Product {
	SQL := "SELECT id, name, price, created_at, updated_at FROM product WHERE id = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	var product domain.Product
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)
	}

	return product
}

func (repository ProductRepositoryImpl) Create(ctx context.Context, db *sql.DB, product domain.Product) domain.Product {
	SQL := "INSERT INTO product(name, price, created_at, updated_at) VALUES(?, ?, ?, ?)"

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := db.ExecContext(ctx, SQL, product.Name, product.Price, product.CreatedAt, product.UpdatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = id
	return product
}

func (repository ProductRepositoryImpl) Update(ctx context.Context, db *sql.DB, product domain.Product) domain.Product {
	SQL := "UPDATE product SET name = ?, price = ?, updated_at = ? WHERE id = ?"

	product.UpdatedAt = time.Now()

	_, err := db.ExecContext(ctx, SQL, product.Name, product.Price, product.UpdatedAt, product.Id)
	helper.PanicIfError(err)

	return product 
}

func (repository ProductRepositoryImpl) Delete(ctx context.Context, db *sql.DB, productId int) {
	SQL := "DELETE FROM product WHERE id = ?"
	_, err := db.ExecContext(ctx, SQL, productId)
	helper.PanicIfError(err)
}
