package router

import (
	"bookstore/db"
	"bookstore/handler"
	"bookstore/middleware"
	repo "bookstore/repository/repo_impl"
	"github.com/labstack/echo"
)

func ProductRouter(e *echo.Echo, sql *db.Sql) {
	handler := handler.ProductHandler{
		ProductRepo: repo.NewProductRepo(sql),
	}

	p := e.Group("/product")
	p.Use(middleware.JWTMiddleware())

	p.POST("/add", handler.Add)
	p.DELETE("/delete/:product_id", handler.Delete)
	p.PUT("/update", handler.Update)
	p.GET("/detail/:product_id", handler.Details)
	p.GET("/list", handler.List)
}
