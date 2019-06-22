package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tiancai110a/gin-blog/models"
	"github.com/tiancai110a/gin-blog/pkg/errno"
	"github.com/tiancai110a/gin-blog/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var errnumber *errno.Errno
	data := make(map[string]interface{})

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": data,
		})
	}()
	username := c.Query("username")
	password := c.Query("password")
	a := auth{Username: username, Password: password}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&a)
	if !ok {
		for _, err := range valid.Errors {
			glog.Error(err.Key, err.Message)
		}
		errnumber = errno.ErrorAuth
		return
	}

	if isExist, err := models.CheckAuth(username, password); !isExist || err != nil {
		errnumber = errno.ErrorAuth
		return
	}
	token, err := util.GenerateToken(username, password)
	if err != nil {
		errnumber = errno.ErrorAuthCheckTokenFail
		return
	}
	errnumber = errno.Success
	data["token"] = token

}
