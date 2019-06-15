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

//获取单个文章
func GetArticle(c *gin.Context) {
	var errnumber *errno.Errno
	valid := validation.Validation{}

	var data models.Article
	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": data,
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
	data = models.GetArticle(int(id))
	errnumber = errno.Success

}

//获取多个文章
func GetArticles(c *gin.Context) {
	valid := validation.Validation{}
	var errnumber *errno.Errno
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": data,
		})
	}()

	state, err := util.ParseAndValidState(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	tagid, err := util.ParseAndValidTagId(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		return
	}

	maps["tag_id"] = tagid
	maps["state"] = state

	errnumber = errno.Success
	data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)
}

//新增文章
func AddArticle(c *gin.Context) {
	valid := validation.Validation{}
	var errnumber *errno.Errno

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": make(map[string]string),
		})
	}()

	state, err := util.ParseAndValidState(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	tagId, err := util.ParseAndValidTagId(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	data := make(map[string]interface{})
	title := util.CheckStringRequired(c, &valid, "title")
	desc := util.CheckStringRequired(c, &valid, "desc")
	content := util.CheckStringRequired(c, &valid, "content")
	createdBy := util.CheckStringRequired(c, &valid, "created_by")

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		return
	}

	if !models.ExistTagById(tagId) {
		errnumber = errno.ErrorNotexistTag
	}

	data["tag_id"] = tagId
	data["title"] = title
	data["desc"] = desc
	data["content"] = content
	data["created_by"] = createdBy
	data["state"] = state

	models.AddArticle(data)

}

//修改文章
func EditArticle(c *gin.Context) {
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

	tagId, err := util.ParseAndValidTagId(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	valid.MaxSize(title, 100, "title").Message("title length up to 100")
	valid.MaxSize(desc, 255, "desc").Message("desc length up to 255")
	valid.MaxSize(content, 65535, "content").Message("desc length up to 100 65535")
	valid.Required(modifiedBy, "modified_by").Message("modifiedBy not exist")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("modifiedBy length up to 100")

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		return
	}

	if !models.ExistArticleById(id) {
		errnumber = errno.ErrorNotexistArticle
		return
	}
	if !models.ExistTagById(tagId) {
		errnumber = errno.ErrorExistTag
		return
	}

	data := make(map[string]interface{})
	if tagId > 0 {
		data["tag_id"] = tagId
	}
	if title != "" {
		data["title"] = title
	}
	if desc != "" {
		data["desc"] = desc
	}
	if content != "" {
		data["content"] = content
	}

	data["modified_by"] = modifiedBy

	models.EditArticle(id, data)

	errnumber = errno.Success
	return

}

//删除文章
func DeleteArticle(c *gin.Context) {
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

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		return
	}
	if !models.ExistArticleById(id) {
		errnumber = errno.Success
		return
	}
	models.DeleteArticle(id)
	errnumber = errno.InvalidParams
	return

}
