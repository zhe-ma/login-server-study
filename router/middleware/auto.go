package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhe-ma/login-server-study/handler"
	"github.com/zhe-ma/login-server-study/pkg/errno"
	"github.com/zhe-ma/login-server-study/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if len(tokenString) == 0 {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		if _, err := token.Parse(tokenString); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
