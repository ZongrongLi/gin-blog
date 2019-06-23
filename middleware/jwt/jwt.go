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
		errnumber := errno.Success
		status := http.StatusOK
		var data interface{}

		errnumber = errno.Success
		token := c.Query("token")
		if token == "" {
			errnumber = errno.ErrorAuth
			status = http.StatusUnauthorized
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				status = http.StatusUnauthorized
				errnumber = errno.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				errnumber = errno.ErrorAuthCheckTokenTimeOut
				status = http.StatusUnauthorized
			}
		}

		if errnumber != errno.Success {
			status = http.StatusUnauthorized

			c.JSON(status, gin.H{
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
