package repository

import (
	"bookstore/model"
	"context"
)

type ProductRepository interface {
	AddProduct(context context.Context, product model.Product) (model.Product, error)
	UpdateProduct(context context.Context, product model.Product) error
	DeleteProduct(context context.Context, product model.Product) error
	SelectProductById(context context.Context, productId string) (model.Product, error)
	SelectAll(context context.Context) ([]model.Product, error)
}
