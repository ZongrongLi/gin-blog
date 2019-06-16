package util

import (
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tiancai110a/gin-blog/pkg/errno"
)

//所有的都是 有才验证 没有就不验证 直接返回成功

func ParseAndValidId(c *gin.Context, valid *validation.Validation) (int, error) {

	idStr := c.Param("id")
	if idStr == "" {
		return -1, nil
	}
	id, err := strconv.ParseInt(idStr, 10, 64)

	valid.Min(id, 1, "id").Message("ID must >0")
	if err != nil {

		glog.Error("state convert failed! err:", err)
		return 0, errno.InvalidParams
	}
	return int(id), nil
}

func ParseAndValidTagId(c *gin.Context, valid *validation.Validation) (int, error) {
	var errnumber *errno.Errno

	idStr := c.Query("tag_id")
	if idStr == "" {
		return -1, nil
	}
	tagid, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		errnumber = errno.InvalidParams
		glog.Error("state convert failed! err:", err)
		return 0, errnumber
	}

	valid.Min(tagid, 1, "tag_id").Message("ID must >0")

	return int(tagid), nil
}
func ParseAndValidState(c *gin.Context, valid *validation.Validation) (int, error) {

	stateStr := c.Query("state")
	if stateStr == "" {
		return -1, nil
	}
	state, err := strconv.ParseInt(stateStr, 10, 64)

	if err != nil {
		glog.Error("state convert failed! err:", err)
		return 0, errno.InvalidParams
	}

	valid.Range(state, 0, 1, "state").Message("state out of range")

	return int(state), nil
}

func ParseAndValidString(c *gin.Context, key string, valid *validation.Validation, limit int) (string, error) {
	value := c.Query(key)
	valid.Required(value, key).Message(key, "not exist")
	valid.MaxSize(value, limit, key).Message(key, "length not more than  100")
	return value, nil
}

func CheckError(valid *validation.Validation) (err *errno.Errno) {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			glog.Errorf("%s %s %s", "param valide failed ：", err.Key, err.Message)
		}
		err = errno.InvalidParams
		return
	}
	err = errno.Success
	return
}

func CheckStringRequired(c *gin.Context, valid *validation.Validation, key string) string {
	value := c.Query(key)
	valid.Required(value, key).Message(key, "not exist")
	return value
}
