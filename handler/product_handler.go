package handler

import (
	"bookstore/encrypt"
	"bookstore/helper"
	"bookstore/model"
	"bookstore/repository"
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type ProductHandler struct {
	ProductRepo repository.ProductRepository
}

func (p *ProductHandler) Add(c echo.Context) error {
	req := model.Product{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	req.ProductId = encrypt.UUIDV1()
	req.UserId = claims.UserId

	product, err := p.ProductRepo.AddProduct(ctx, req)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	fmt.Println(product)

	return helper.ResponseData(c, product)
}

func (p *ProductHandler) Delete(c echo.Context) error {
	defer c.Request().Body.Close()

	productId := c.Param("product_id")
	if len(productId) == 0 {
		return helper.ResponseErr(c, http.StatusBadRequest)
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	product := model.Product{
		ProductId: productId,
		UserId:    claims.UserId,
	}
	err := p.ProductRepo.DeleteProduct(ctx, product)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}
	return helper.ResponseData(c, "Delete thành công")
}

func (p *ProductHandler) Update(c echo.Context) error {
	req := model.Product{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)
	req.UserId = claims.UserId

	err := p.ProductRepo.UpdateProduct(ctx, req)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, "Update thành công")
}

func (p *ProductHandler) Details(c echo.Context) error {
	defer c.Request().Body.Close()

	productId := c.Param("product_id")
	if len(productId) == 0 {
		return helper.ResponseErr(c, http.StatusBadRequest, "Thiếu id sản phẩm")
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	product, err := p.ProductRepo.SelectProductById(ctx, productId)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	if product == (model.Product{}) {
		return helper.ResponseErr(c, http.StatusNotFound, "Sản phẩm này không tồn tại")
	}

	if product.DeletedAt.Valid {
		return helper.ResponseErr(c, http.StatusNotFound, "Sản phẩm này Đã bị xoá")
	}

	product.UserId = ""
	return helper.ResponseData(c, product)
}

func (p *ProductHandler) List(c echo.Context) error {
	defer c.Request().Body.Close()

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	products, err := p.ProductRepo.SelectAll(ctx)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, products)
}
