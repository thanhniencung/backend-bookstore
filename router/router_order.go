package router

import (
	"bookstore/db"
	"bookstore/middleware"
	"bookstore/handler"
	repo "bookstore/repository/repo_impl"
	"github.com/labstack/echo"
)

func OrderRouter(e *echo.Echo, sql *db.Sql) {
	handler := handler.OrderHandler{
		OrderRepo: repo.NewOrderRepo(sql),
	}

	c := e.Group("/order")
	c.Use(middleware.JWTMiddleware())

	c.POST("/add", handler.AddToCard)
	c.POST("/confirm", handler.Confirm)
	c.POST("/update", handler.Update)
	c.GET("/count", handler.OrderCountItem)
	c.GET("/detail", handler.OrderDetails)
	c.GET("/list", handler.OrderList)
}
