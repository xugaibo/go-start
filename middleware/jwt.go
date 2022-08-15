package middleware

import (
	"github.com/gin-gonic/gin"
	"go-start/core/bizcode"
	"go-start/core/jwtutil"
	"go-start/models/response"
	"net/http"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		jwtToken, err := jwtutil.ParseToken(token)
		if err == nil && jwtToken != nil {
			if jwtToken.IsExpire() {
				c.AbortWithStatusJSON(http.StatusOK, response.Biz(bizcode.TokenExpire))
				return
			}
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusOK, response.Biz(bizcode.TokenInvalid))
		}
	}
}
