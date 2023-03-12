package middlewares

import (
	"alta-airbnb-be/app/config"
	"alta-airbnb-be/utils/consts"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

//go:generate mockery --name JWTMiddleware_ --output ../mocks
type JWTMiddleware_ interface {
	ExtractToken(e echo.Context) (uint, string, error)
}

type JWT struct{}

// ExtractToken implements JWTMiddleware_
func (u *JWT) ExtractToken(e echo.Context) (uint, string, error) {
	return ExtractToken(e)
}

func NewJWT() JWTMiddleware_ {
	return &JWT{}
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echoJwt.WithConfig(echoJwt.Config{
		SigningKey:    []byte(config.SECRET_JWT),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId int, userRole string) (string, error) {
	claims := jwt.MapClaims{}
	claims[consts.JWT_Authorized] = true
	claims[consts.JWT_UserId] = userId
	claims[consts.JWT_Role] = userRole
	claims[consts.JWT_ExpiredTime] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractToken(e echo.Context) (uint, string, error) {
	token, ok := e.Get("user").(*jwt.Token)
	if !ok {
		return 0, "", errors.New(consts.JWT_InvalidJwtToken)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New(consts.JWT_FailedCastingJwtToken)
	}
	if token.Valid {
		loggedInUserId, existedUserId := claims[consts.JWT_UserId].(float64)
		loggedInUserRole, existedUserRole := claims[consts.JWT_Role]
		if !existedUserId || !existedUserRole {
			return 0, "", errors.New(consts.SERVER_InternalServerError)
		}
		return uint(loggedInUserId), loggedInUserRole.(string), nil
	}
	return 0, "", errors.New(consts.SERVER_InternalServerError)
}
