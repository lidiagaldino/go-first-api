package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/authentication"
	"github.com/lidiagaldino/go-first-api/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			err := authentication.VerifyToken(authToken)
			if err == nil {
				c.Next()
				return
			}
			utils.SendError(c, http.StatusUnauthorized, "Not authorized")
			c.Abort()
			return
		}
		utils.SendError(c, http.StatusUnauthorized, "Not authorized")
		c.Abort()
	}
}
