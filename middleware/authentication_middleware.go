package middleware

import (
	"errors"
	"fmt"
	"genesis/configs"
	"genesis/responses"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email          string `json:"email"`
	OrganizationId string `json:"organizationId"`
	jwt.StandardClaims
}

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, responses.ContributionResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    "Missing required authorization header",
			})
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(token, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, responses.ContributionResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthroized",
				Data:    "Invalid token",
			})
			ctx.Abort()
			return
		}

		token = tokenParts[1]

		claims, err := ValidateToken(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, responses.ContributionResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauathorized",
				Data:    "Invalid credentials",
			})
			ctx.Abort()
			return
		}

		ctx.Set("organizationId", claims.OrganizationId)
		ctx.Next()
	}
}

func ValidateToken(tokenString string) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signature")
		}

		return []byte(configs.AppConfig.JWT_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claims)

	if !ok || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
