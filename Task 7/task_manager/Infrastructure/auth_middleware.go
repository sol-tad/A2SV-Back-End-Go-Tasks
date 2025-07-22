package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("solomon1234")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader:=c.GetHeader("Authorization")
		if authHeader==""||!strings.HasPrefix(authHeader,"Bearer"){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}
		tokenStr:=strings.TrimPrefix(authHeader,"Bearer")


		token,err:=jwt.Parse(tokenStr,func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		claims,_:=token.Claims.(jwt.MapClaims)

		
		c.Set("username",claims["username"])
		c.Set("role",claims["role"])

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role,exists:=c.Get("role")
		if!exists||role!="admin"{
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			return
		}
		c.Next()

	}
}