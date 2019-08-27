package main

import (
	"bookstore/db"
	"bookstore/model"
	"bookstore/router"
	"github.com/labstack/echo"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "ryannguyen",
		Password: "123456",
		DbName:   "bookstore",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, model.Response{
			StatusCode: 200,
			Message:    "Home Page",
		})
	})

	router.UserRouter(e, sql)
	router.CateRouter(e, sql)
	router.ProductRouter(e, sql)
	router.OrderRouter(e, sql)

	e.Logger.Fatal(e.Start(":8000"))
}
