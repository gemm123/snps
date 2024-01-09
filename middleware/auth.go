package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"synapsis/internal/model"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"meta": model.Meta{
					Message: "unauthorized",
					Status:  401,
				},
			})
			ctx.Abort()
			return
		}

		if !strings.Contains(auth, "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"meta": model.Meta{
					Message: "invalid authorization",
					Status:  401,
				},
			})
			ctx.Abort()
			return
		}

		var tokenString string
		arrayAuth := strings.Split(auth, " ")
		if len(arrayAuth) == 2 {
			tokenString = arrayAuth[1]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("invalid token")
			}

			return []byte(os.Getenv("JWT_SECRETKEY")), nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"meta": model.Meta{
					Message: "invalid token",
					Status:  401,
				},
			})
			ctx.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		ctx.Set("id", claims["id"])
		ctx.Set("email", claims["email"])

		ctx.Next()
	}
}
