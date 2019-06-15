package v1

import (
	"net/http"
	"strconv"

	"github.com/golang/glog"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/tiancai110a/gin-blog/models"
	"github.com/tiancai110a/gin-blog/pkg/errno"
	"github.com/tiancai110a/gin-blog/pkg/setting"
	"github.com/tiancai110a/gin-blog/util"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	state := c.Query("state")
	if state != "" {
		maps["state"] = state
	}
	tags := models.GetTags(util.GetPage(c), setting.PageSize, maps)
	cnt := models.GetTagTotal(maps)
	data["list"] = tags
	data["count"] = cnt
	c.JSON(http.StatusOK, gin.H{
		"code": errno.Success.Code,
		"msg":  errno.Success.Message,
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	var errnumber *errno.Errno

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": make(map[string]string),
		})
	}()
	state, err := strconv.ParseInt(c.Query("state"), 10, 64)
	if err != nil {
		errnumber = errno.InvalidParams
		glog.Error("state convert failed! err:", err, c.Query("state"))
		return
	}
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name not exist")
	valid.MaxSize(name, 100, "name").Message("name length not more than  100")
	valid.Required(createdBy, "created_by").Message("createdBy not exist")
	valid.MaxSize(createdBy, 100, "name").Message("createdBy length not more than  100")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			glog.Errorf("%s %s %s", "param valide failed ：", err.Key, err.Message)
		}
		errnumber = errno.InvalidParams
		return
	}
	if models.ExistTagByName(name) {
		errnumber = errno.ErrorExistTag
		return
	}
	models.AddTag(name, int(state), createdBy)
	errnumber = errno.Success

}

//修改文章标签
func EditTag(c *gin.Context) {
	var errnumber *errno.Errno

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": make(map[string]string),
		})
	}()
	id := c.Param("id")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{}
	name := c.Query("name")
	state := c.Query("state")

	istate, err := strconv.ParseInt(state, 10, 64)
	if err != nil {
		glog.Info("convert id failed err:", err)
	}
	if state != "" {
		valid.Range(istate, 0, 1, "state").Message("state out of range")
	}
	valid.Required(id, "id").Message("id is not exist")
	valid.Required(name, "name").Message("name not exist")
	valid.MaxSize(name, 100, "name").Message("name length not more than  100")
	valid.Required(modifiedBy, "modified_by").Message("modified_by not exist")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("modifiedBy length not more than  100")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			glog.Errorf("%s %s %s", "param valide failed ：", err.Key, err.Message)
		}
		errnumber = errno.InvalidParams
		return
	}
	if !models.ExistTagById(id) {
		errnumber = errno.ErrorNotexistTag
		return
	}

	errnumber = errno.Success

	models.EditTag(id, name, int(istate), modifiedBy)

}

//删除文章标签
func DeleteTag(c *gin.Context) {
	var errnumber *errno.Errno
	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": make(map[string]string),
		})
	}()

	id := c.Param("id")
	id_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		glog.Info("convert id failed err:", err)
	}
	valid := validation.Validation{}
	valid.Min(id_, 1, "id").Message("ID must >0")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			glog.Errorf("%s %s %s", "param valide failed ：", err.Key, err.Message)
		}
		errnumber = errno.InvalidParams
		return
	}
	errnumber = errno.Success
	if models.ExistTagById(id) {
		models.DeleteTag(id)
	}
}
