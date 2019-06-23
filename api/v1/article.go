package v1

import (
	//	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tiancai110a/gin-blog/models"
	"github.com/tiancai110a/gin-blog/service/article_service"
	"github.com/tiancai110a/gin-blog/service/tag_service"
	"net/http"

	"github.com/tiancai110a/gin-blog/pkg/app"
	"github.com/tiancai110a/gin-blog/pkg/errno"
	"github.com/tiancai110a/gin-blog/pkg/qrcode"
	"github.com/tiancai110a/gin-blog/pkg/setting"
	"github.com/tiancai110a/gin-blog/util"
)

const (
	QRCODE_URL = "https://github.com/tiancai110a/gin-blog/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
)

// @Summary 获取单个文章
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/:id [get]

func GetArticle(c *gin.Context) {
	errnumber := errno.Success
	valid := validation.Validation{}

	var data *models.Article
	appG := app.Gin{c}
	defer func() {
		appG.Response(http.StatusOK, errnumber, data)
	}()

	id, err := util.ParseAndValidId(c, &valid)
	if err != nil {
		glog.Error("id valid fail")

		errnumber = errno.InvalidParams
		return
	}

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		glog.Error("validate  error ")
		return
	}

	cache := article_service.Article{Id: id}

	data, err = cache.Get()
	if err != nil {
		glog.Error("get article failed", err)
		errnumber = errno.ErrorInternel
		return
	}
}

// @Summary 获取多个文章
// @Produce  json
// @Param id query string true "id"
// @Param tagid query string true "tagid"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [get]

func GetArticles(c *gin.Context) {
	valid := validation.Validation{}
	errnumber := errno.Success
	data := make(map[string]interface{})

	appG := app.Gin{c}
	defer func() {
		appG.Response(http.StatusOK, errnumber, data)
	}()

	cache := article_service.Article{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	state, err := util.ParseAndValidState(c, &valid)
	if err != nil {
		glog.Error("state valid fail")
		errnumber = errno.InvalidParams
		return
	}
	if state >= 0 {
		cache.State = state
	}

	tagid, err := util.ParseAndValidTagId(c, &valid)
	if err != nil {
		glog.Error("tagid valid fail")
		errnumber = errno.InvalidParams
		return
	}
	if tagid >= 0 {
		cache.TagId = tagid
	}

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		glog.Error("validate  error ")
		return
	}

	errnumber = errno.Success

	data["list"], err = cache.GetAll()
	if err != nil {
		glog.Error("cache.GetAll() err:", err)
		errnumber = errno.ErrorGetArticleFailed
		return
	}
	data["total"], err = cache.Count()
	if err != nil {
		glog.Error("get articles failed", err)
		errnumber = errno.ErrorInternel
		return
	}
}

// data := make(map[string]interface{})
// 	title := util.CheckStringRequired(c, &valid, "title")
// 	desc := util.CheckStringRequired(c, &valid, "desc")
// 	content := util.CheckStringRequired(c, &valid, "content")
// 	createdBy := util.CheckStringRequired(c, &valid, "created_by")
// @Summary 获取多个文章
// @Produce  json
// @Param id query string true "id"
// @Param tagid query string true "tagid"
// @Param data query string true "data"
// @Param title query string true "title"
// @Param desc query string true "desc"
// @Param content query string true "content"
// @Param createdBy query string true "createdBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
	valid := validation.Validation{}
	errnumber := errno.Success

	appG := app.Gin{c}
	defer func() {
		appG.Response(http.StatusOK, errnumber, make(map[string]string))
	}()

	state, err := util.ParseAndValidState(c, &valid)
	if err != nil {
		glog.Error("state valid fail")
		errnumber = errno.InvalidParams
		return
	}

	tagId, err := util.ParseAndValidTagId(c, &valid)
	if err != nil {
		glog.Error("tagId valid fail")
		errnumber = errno.InvalidParams
		return
	}

	title := util.CheckStringRequired(c, &valid, "title")
	desc := util.CheckStringRequired(c, &valid, "desc")
	content := util.CheckStringRequired(c, &valid, "content")
	createdBy := util.CheckStringRequired(c, &valid, "created_by")
	coverimageurl := c.Query("cover_image_url")

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		glog.Error("validate  error ")
		return
	}
	tagService := tag_service.Tag{Id: tagId}

	if !tagService.ExistById() {
		errnumber = errno.ErrorGetExistTagFail
		return
	}

	cache := article_service.Article{
		TagId:         tagId,
		Title:         title,
		Desc:          desc,
		Content:       content,
		CreatedBy:     createdBy,
		State:         state,
		CoverImageUrl: coverimageurl,
	}
	err = cache.Add()
	if err != nil {
		glog.Error("add article failed", err)
		errnumber = errno.ErrorInternel
		return
	}
}

// data := make(map[string]interface{})
// 	title := util.CheckStringRequired(c, &valid, "title")
// 	desc := util.CheckStringRequired(c, &valid, "desc")
// 	content := util.CheckStringRequired(c, &valid, "content")
// 	createdBy := util.CheckStringRequired(c, &valid, "created_by")
// @Summary 获取多个文章
// @Produce  json
// @Param id query string true "id"
// @Param tagid query string true "tagid"
// @Param data query string true "data"
// @Param title query string true "title"
// @Param desc query string true "desc"
// @Param content query string true "content"
// @Param createdBy query string true "createdBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [put]
func EditArticle(c *gin.Context) {
	var errnumber *errno.Errno
	valid := validation.Validation{}

	appG := app.Gin{c}
	defer func() {
		appG.Response(http.StatusOK, errnumber, make(map[string]string))
	}()
	id, err := util.ParseAndValidId(c, &valid)
	if err != nil {
		errnumber = errno.InvalidParams
		return
	}

	tagId, err := util.ParseAndValidTagId(c, &valid)
	if err != nil {
		glog.Error("InvalidParams")
		errnumber = errno.InvalidParams
		return
	}

	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	coverimageurl := c.Query("cover_image_url")

	valid.MaxSize(title, 100, "title").Message("title length up to 100")
	valid.MaxSize(desc, 255, "desc").Message("desc length up to 255")
	valid.MaxSize(content, 65535, "content").Message("desc length up to 100 65535")
	valid.Required(modifiedBy, "modified_by").Message("modifiedBy not exist")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("modifiedBy length up to 100")

	errnumber = util.CheckError(&valid)
	if errnumber != errno.Success {
		return
	}
	cache := article_service.Article{
		Id:            id,
		Title:         title,
		Desc:          desc,
		Content:       content,
		ModifiedBy:    modifiedBy,
		CoverImageUrl: coverimageurl,
	}
	if !cache.ExistById() {
		glog.Error("err exist artcle:", id)
		errnumber = errno.ErrorNotExistArticle
		return
	}

	if tagId > 0 {
		cache.TagId = tagId
		tagService := tag_service.Tag{Id: tagId}
		if !tagService.ExistById() {
			glog.Error("err exist tag:", tagId)
			errnumber = errno.ErrorExistTag
			return
		}
	}
	if title != "" {
		cache.Title = title
	}
	if desc != "" {
		cache.Desc = desc
	}
	if content != "" {
		cache.Content = content
	}

	cache.ModifiedBy = modifiedBy

	err = cache.Add()
	if err != nil {
		errnumber = errno.ErrorInternel
		glog.Error("add failed err:", err)
		return
	}
	errnumber = errno.Success
	return

}

// @Summary 获取多个文章
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [delete]

//删除文章
func DeleteArticle(c *gin.Context) {
	errnumber := errno.Success
	valid := validation.Validation{}

	appG := app.Gin{c}
	defer func() {
		appG.Response(http.StatusOK, errnumber, make(map[string]string))
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

	articleService := article_service.Article{Id: id}

	if !articleService.ExistById() {
		errnumber = errno.ErrorNotExistArticle
		return
	}

	articleService.Delete()
	errnumber = errno.InvalidParams
	return

}

func GenerateArticlePoster(c *gin.Context) {
	appG := app.Gin{c}
	article := &article_service.Article{}

	qr := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto) // 目前写死 gin 系列路径，可自行增加业务逻辑
	posterName := article_service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
	articlePoster := article_service.NewArticlePoster(posterName, article, qr)
	articlePosterBgService := article_service.NewArticlePosterBg(
		"bg.jpg",
		articlePoster,
		&article_service.Rect{
			X0: 0,
			Y0: 0,
			X1: 550,
			Y1: 700,
		},
		&article_service.Pt{
			X: 125,
			Y: 298,
		},
	)

	_, filePath, err := articlePosterBgService.Generate()
	if err != nil {
		appG.Response(http.StatusOK, errno.ErrorGenArticlePosterFail, nil)
		return
	}

	appG.Response(http.StatusOK, errno.Success, map[string]string{
		"poster_url":      qrcode.GetQrCodeFullUrl(posterName),
		"poster_save_url": filePath + posterName,
	})
}
