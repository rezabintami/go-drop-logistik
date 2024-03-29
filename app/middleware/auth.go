package middleware

import (
	"fmt"
	"go-drop-logistik/helpers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT        string
	RefreshSecretJWT string
	ExpiresDuration  int
}

type JwtRefreshToken struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigMiddleware struct {
	logger helpers.Logger
}

func NewMiddleware(logger helpers.Logger) ConfigMiddleware {
	return ConfigMiddleware{
		logger: logger,
	}
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if result, ok := err.(*echo.HTTPError); ok {
		helpers.ErrorResponse(c, result.Code, fmt.Errorf("%v", result.Message))
	} else {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
}

// GenerateToken jwt ...
func (jwtConf *ConfigJWT) GenerateToken(userID int, name, role string) string {
	claims := JwtCustomClaims{
		userID,
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token
}

// GenerateToken jwt ...
func (jwtConf *ConfigJWT) GenerateRefreshToken(userID int) string {
	claims := JwtRefreshToken{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.RefreshSecretJWT))

	return token
}

// GetUser from jwt ...
func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

func (logs *ConfigMiddleware) MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logs.logger.Logging(c).Info("incoming request")
		return next(c)
	}
}

// func MiddlewareLoggingEntry(c echo.Context, request, response interface{}) {
// 	logging.LogEntry(c, request, response).Info("incoming request")
// }
