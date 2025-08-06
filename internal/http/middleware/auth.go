package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("change-me")

// AuthRequired valida el token y coloca el usuario en el contexto.
func AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(auth, "Bearer ")

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	uid, ok := claims["user_id"].(float64)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("currentUserID", uint(uid))
	c.Next()
}
