package router

import (
	"bookstore/db"
	"bookstore/middleware"
	"bookstore/handler"
	repo "bookstore/repository/repo_impl"
	"github.com/labstack/echo"
)

func UserRouter(e *echo.Echo, sql *db.Sql) {

	handler := handler.UserHandler{
		UserRepo: repo.NewUserRepo(sql),
	}

	e.POST("/user/sign-in", handler.SignIn)
	e.POST("/user/sign-up", handler.SignUp)

	e.GET("/user/profile", handler.Profile, middleware.JWTMiddleware())
	e.GET("/user/list", handler.List, middleware.JWTMiddleware())
}
