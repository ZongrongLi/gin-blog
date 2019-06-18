package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/tiancai110a/gin-blog/models"
	"github.com/tiancai110a/gin-blog/pkg/errno"
	"github.com/tiancai110a/gin-blog/pkg/setting"
	"github.com/tiancai110a/gin-blog/util"
)

// @Summary 获取多个文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	var errnumber *errno.Errno
	valid := validation.Validation{}

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": data,
		})
	}()
	name, err := util.ParseAndValidString(c, "name", &valid, 100)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	state, err := util.ParseAndValidState(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	maps["name"] = name
	maps["state"] = state

	errnumber = util.CheckError(&valid)

	if errnumber != errno.Success {
		return
	}

	tags := models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	cnt := models.GetTagTotal(maps)
	data["list"] = tags
	data["count"] = cnt
	errnumber = errno.Success
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	valid := validation.Validation{}
	var errnumber *errno.Errno

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": make(map[string]string),
		})
	}()

	name, err := util.ParseAndValidString(c, "name", &valid, 100)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	state, err := util.ParseAndValidState(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	createdBy, err := util.ParseAndValidString(c, "created_by", &valid, 100)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		return
	}

	if models.ExistTagByName(name) {
		errnumber = errno.ErrorExistTag
		return
	}
	models.AddTag(name, int(state), createdBy)
	errnumber = errno.Success

}

// @Summary 修改文章标签
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [put]

func EditTag(c *gin.Context) {
	var errnumber *errno.Errno
	valid := validation.Validation{}

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": make(map[string]string),
		})
	}()

	id, err := util.ParseAndValidId(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	name, err := util.ParseAndValidString(c, "name", &valid, 100)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	state, err := util.ParseAndValidState(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}
	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		return
	}

	modifiedBy, err := util.ParseAndValidString(c, "modified_by", &valid, 100)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	if !models.ExistTagById(id) {
		errnumber = errno.ErrorNotexistTag
		return
	}

	errnumber = errno.Success

	models.EditTag(id, name, int(state), modifiedBy)

}

// @Summary 修改文章标签
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [delete]
func DeleteTag(c *gin.Context) {
	valid := validation.Validation{}
	var errnumber *errno.Errno
	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": make(map[string]string),
		})
	}()

	id, err := util.ParseAndValidId(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		return
	}

	errnumber = errno.Success
	if models.ExistTagById(id) {
		models.DeleteTag(id)
	}
}
