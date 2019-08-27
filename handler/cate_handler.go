package handler

import (
	"bookstore/encrypt"
	"bookstore/helper"
	"bookstore/model"
	"bookstore/repository"
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type CateHandler struct {
	CateRepo repository.CateRepository
}

func (m *CateHandler) Add(c echo.Context) error {
	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	if claims.Role != model.ADMIN.String() {
		return helper.ResponseErr(c, http.StatusForbidden, "Lỗi quyền truy cập") // permission denied
	}

	req := model.Cate{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	req.CateId = encrypt.UUIDV1()

	cate, err := m.CateRepo.AddCate(ctx, req)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, cate)
}

func (m *CateHandler) Delete(c echo.Context) error {
	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	if claims.Role != model.ADMIN.String() {
		return helper.ResponseErr(c, http.StatusForbidden, "Lỗi quyền truy cập") // permission denied
	}

	req := model.Cate{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	err := m.CateRepo.DeleteCate(ctx, req.CateId)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, "Delete thành công")
}

func (m *CateHandler) Update(c echo.Context) error {
	// Lấy thông tin user_id từ token
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*model.JwtCustomClaims)

	if claims.Role != model.ADMIN.String() {
		return helper.ResponseErr(c, http.StatusForbidden, "Lỗi quyền truy cập") // permission denied
	}

	req := model.Cate{}
	defer c.Request().Body.Close()

	if err := c.Bind(&req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return helper.ResponseErr(c, http.StatusBadRequest, err.Error())
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)

	err := m.CateRepo.UpdateCate(ctx, req)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, "Update thành công")
}

func (m *CateHandler) Details(c echo.Context) error {
	defer c.Request().Body.Close()

	cateId := c.Param("cate_id")
	if len(cateId) == 0 {
		return helper.ResponseErr(c, http.StatusBadRequest, "Thiếu id danh mục")
	}

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	cate, err := m.CateRepo.SelectCateById(ctx, cateId)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	if cate == (model.Cate{}) {
		return helper.ResponseErr(c, http.StatusNotFound, "Danh mục này không tồn tại")
	}

	if cate.DeletedAt.Valid {
		return helper.ResponseErr(c, http.StatusNotFound, "Danh mục này Đã bị xoá")
	}

	return helper.ResponseData(c, cate)
}

func (m *CateHandler) List(c echo.Context) error {
	defer c.Request().Body.Close()

	ctx, _ := context.WithTimeout(c.Request().Context(), 10*time.Second)
	cates, err := m.CateRepo.SelectAll(ctx)
	if err != nil {
		return helper.ResponseErr(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseData(c, cates)
}
