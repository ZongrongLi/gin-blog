package util

import (
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tiancai110a/gin-blog/pkg/errno"
)

func ParseAndValidId(c *gin.Context, valid *validation.Validation) (int, error) {
	var errnumber *errno.Errno

	idStr := c.Param("id")
	if idStr == "" {
		return 0, errno.InvalidParams
	}
	id, err := strconv.ParseInt(idStr, 10, 64)

	valid.Min(id, 1, "id").Message("ID must >0")
	if err != nil {
		errnumber = errno.InvalidParams
		glog.Error("state convert failed! err:", err)
		return 0, errnumber
	}
	return int(id), nil
}

func ParseAndValidTagId(c *gin.Context, valid *validation.Validation) (int, error) {
	var errnumber *errno.Errno

	idStr := c.Param("tag_id")
	if idStr == "" {
		return 0, errno.InvalidParams
	}
	tagid, err := strconv.ParseInt(idStr, 10, 64)

	valid.Min(tagid, 1, "tag_id").Message("ID must >0")
	if err != nil {
		errnumber = errno.InvalidParams
		glog.Error("state convert failed! err:", err)
		return 0, errnumber
	}
	return int(tagid), nil
}
func ParseAndValidState(c *gin.Context, valid *validation.Validation) (int, error) {
	var errnumber *errno.Errno

	stateStr := c.Query("state")
	if stateStr == "" {
		return 0, errno.Success
	}
	state, err := strconv.ParseInt(stateStr, 10, 64)

	if err != nil {
		errnumber = errno.InvalidParams
		glog.Error("state convert failed! err:", err)
		return 0, errnumber
	}

	valid.Range(state, 0, 1, "state").Message("state out of range")

	return int(state), nil
}

func ParseAndValidString(c *gin.Context, key string, valid *validation.Validation, limit int) (string, error) {
	value := c.Query(key)
	valid.Required(value, key).Message(key, "not exist")
	valid.MaxSize(valid, limit, key).Message(key, "length not more than  100")
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
