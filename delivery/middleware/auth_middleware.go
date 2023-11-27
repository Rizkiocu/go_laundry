package middleware

import (
	"go_laundry/util/security"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//logic Here
		var authHeader authHeader
		if err := c.ShouldBindHeader(&authHeader); err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
			})
			return
		}

		//verifikasi jwt token
		token := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", 1)
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
			})
			return
		}

		// verifikasi
		claims, err := security.VerifyJWTToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
			})
			return
		}
		c.Set("claims", claims)

		c.Next()
	}
}
