package middleware

import (
	"go-drop-logistik/helper/logging"
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
	SecretJWT       string
	ExpiresDuration int
}

type ConfigMiddleware struct {
	logger logging.Logger
}

func NewMiddleware(logger logging.Logger) ConfigMiddleware {
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
