package middleware

import (
	"bookstore/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
)

const SECRET_KEY  = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqGKukO1De7zhZj6+H0qtjTkVxwTCpvKe4eCZ0FPqri0cb2JZfXJ/DgYSF6vUpwmJG8wVQZKjeGcjDOL5UlsuusFncCzWBQ7RKNUSesmQRMSGkVb1/3j+skZ6UtW+5u09lHNsj6tQ51s1SPrCBkedbNf0Tp0GbMJDyR4e9T04ZZwIDAQAB"

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(SECRET_KEY),
	}
	return middleware.JWTWithConfig(config)
}

func GenToken(user model.User) (string, error) {
	if len(user.UserId) == 0 || len(user.Role) == 0 || len(user.Phone) == 0 {
		return "", errors.New("len(UserId) == 0 || len(Role) == 0 || len(Phone) == 0")
	}

	claims := &model.JwtCustomClaims{
		UserId:      user.UserId,
		Role:        user.Role,
		PhoneNumber: user.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3600).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return result, nil
}
