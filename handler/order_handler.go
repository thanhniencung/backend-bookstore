package handler

import (
	"bookstore/helper"
	"bookstore/model"
	"bookstore/repository"
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"fmt"
)

type OrderHandler struct {
	OrderRepo repository.OrderRepository
}

// Click nút mua là call api này
func (o *OrderHandler) AddToCard(c echo.Context) error {
	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	req := model.Card{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)

	cart, err := o.OrderRepo.AddToCard(ctx, claims.UserId, req)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, echo.Map{"total": cart.Total, "orderId": cart.OrderId})
}

// Click vào shopping card icon ở AppBar sẽ call api này
// Lấy toàn bộ thông tin của order
func (o *OrderHandler) OrderDetails(c echo.Context) error {
	orderId := c.QueryParam("order_id")
	if len(orderId) == 0 {
		return helper.ResponseErr(c, http.StatusBadRequest, "Thiếu id đơn hàng")
	}

	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	defer c.Request().Body.Close()

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	result, err := o.OrderRepo.ShoppingCard(ctx, claims.UserId, orderId)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, result)
}

func (o *OrderHandler) Update(c echo.Context) error {
		userData := c.Get("user").(*jwt.Token)
		claims := userData.Claims.(*model.JwtCustomClaims)

		req := model.Card{}
		defer c.Request().Body.Close()

		if err := c.Bind(&req); err != nil {
			return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
		}

		if len(req.OrderId) == 0 || req.Quantity == 0 || len(req.ProductId) == 0 {
			return helper.ResponseErr(c, http.StatusBadRequest)
		}

		ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
		err := o.OrderRepo.UpdateQuantityOrder(ctx, claims.UserId, req.OrderId, req.Quantity, req.ProductId)
		if err != nil {
			return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
		}

		return helper.ResponseData(c, echo.Map{"quantity": req.Quantity})
}

// Xác nhận order này đã chuẩn ở phía user
func (o *OrderHandler) Confirm(c echo.Context) error {
	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	req := model.Order{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	req.UserId = claims.UserId

	fmt.Println(req)

	err := o.OrderRepo.UpdateStateOrder(ctx, req)
	if err != nil {
		return helper.ResponseErr(c, http.StatusNotFound, err.Error())
	}

	return helper.ResponseData(c, nil)
}

// Khi vào app gọi api này để hiển thị số lượng sản phẩm có trong shopping card
// Hiển thị ở phần icon AppBar
func (o *OrderHandler) OrderCountItem(c echo.Context) error {
	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	defer c.Request().Body.Close()

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	result, err := o.OrderRepo.CountShoppingCard(ctx, claims.UserId)
	
	if err != nil {
		if result.Total == -1 {
			return helper.ResponseErr(c, http.StatusNotFound, err.Error())
		}
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}
	
	return helper.ResponseData(c, result)
}

func (o *OrderHandler) OrderList(c echo.Context) error {
	defer c.Request().Body.Close()

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	result, err := o.OrderRepo.ListOrder(ctx)

	if err != nil {
		return helper.ResponseErr(c, http.StatusNotFound, err.Error())
	}

	return helper.ResponseData(c, result)
}


