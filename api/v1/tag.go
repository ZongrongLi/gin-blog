package v1

import (
	"github.com/golang/glog"
	"github.com/tiancai110a/gin-blog/service/tag_service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/tiancai110a/gin-blog/models"
	"github.com/tiancai110a/gin-blog/pkg/app"
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
	var errnumber *errno.Errno
	valid := validation.Validation{}

	appG := app.Gin{c}
	defer func() {
		appG.Response(http.StatusOK, errnumber, data)
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

	errnumber = util.CheckError(&valid)

	if errnumber != errno.Success {
		return
	}

	cache := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	tags, err := cache.Get()
	if err != nil {
		glog.Error("tags get failed err:", err)
		errnumber = errno.ErrorInternel
		return
	}
	cnt, err := cache.Count()
	if err != nil {
		glog.Error("tags count get failed err:", err)
		errnumber = errno.ErrorInternel
		return
	}
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
	appG := app.Gin{c}

	defer func() {
		appG.Response(http.StatusOK, errnumber, make(map[string]string))
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

	cache := tag_service.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}

	if cache.ExistByName() {
		errnumber = errno.ErrorExistTag
		return
	}
	err = cache.Add()
	if err != nil {
		errnumber = errno.ErrorInternel
		return
	}
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

	cache := tag_service.Tag{
		Id:         id,
		Name:       name,
		State:      state,
		ModifiedBy: modifiedBy,
	}
	if cache.ExistById() {
		errnumber = errno.ErrorNotExistTag
		return
	}
	cache.Edit()
	errnumber = errno.Success
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
