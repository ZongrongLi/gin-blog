package jwt

import (
	"net/http"
	"time"

	"github.com/tiancai110a/gin-blog/pkg/errno"
	"github.com/tiancai110a/gin-blog/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errnumber *errno.Errno
		var data interface{}

		errnumber = errno.Success
		token := c.Query("token")
		if token == "" {
			errnumber = errno.InvalidParams
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				errnumber = errno.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				errnumber = errno.ErrorAuthCheckTokenTimeOut
			}
		}

		if errnumber != errno.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": errnumber.Code,
				"msg":  errnumber.Message,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
