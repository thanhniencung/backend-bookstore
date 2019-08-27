package repository

import (
	"bookstore/model"
	"context"
)

type OrderRepository interface {
	UpdateStateOrder(context context.Context, order model.Order) error
	UpdateQuantityOrder(context context.Context, userId string, orderId string, quantity int, productId string) error
	AddToCard(context context.Context, userId string, card model.Card) (int, error)
	CountShoppingCard(context context.Context, userId string) (model.OrderCount, error)
	ShoppingCard(context context.Context, userId string, orderId string) (model.Order, error)
	ListOrder(context context.Context) ([]model.Order, error)
}
