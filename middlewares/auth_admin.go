package middlewares

import (
	"mall/tool"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("token")
		if auth == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无登录凭证，请先登录",
			})
			c.Abort()
			return
		}
		_, err := tool.ParseToken(auth)
		if err != nil {
			if err == jwt.ErrTokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "登录凭证已过期，请重新登录",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "错误的登录凭证，请重新登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
