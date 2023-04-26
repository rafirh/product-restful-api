package repository

import (
	"context"
	"database/sql"
	"errors"
	"product_restful_api/helper"
	"product_restful_api/model/domain"
	"time"
)

type ProductRepositoryImpl struct {}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	sql := "INSERT INTO product(name, price, created_at, updated_at) VALUES(?, ?)"

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := tx.ExecContext(ctx, sql, product.Name, product.Price)
	helper.PanicIfError(err)
	
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	product.Id = int(id)

	return product
}

func (repository ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	sql := "UPDATE product SET name = ?, price = ?, updated_at = ? WHERE id = ?"
	product.UpdatedAt = time.Now()
	_, err := tx.ExecContext(ctx, sql, product.Name, product.Price, product.UpdatedAt, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	sql := "DELETE FROM product WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, product.Id)
	helper.PanicIfError(err)
}

func (repository ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	sql := "SELECT id, name, price, created_at, updated_at FROM product WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, sql, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	var product domain.Product
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)
	} else {
		return product, errors.New("product is not found")
	}
	
	return product, nil
}

func (repository ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	sql := "SELECT id, name, price, created_at, updated_at FROM product"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	var products []domain.Product
	if rows.Next() {
		var product domain.Product
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		products = append(products, product)
	}

	return products
}