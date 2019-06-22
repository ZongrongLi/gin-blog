package article_service

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/tiancai110a/gin-blog/models"
	"github.com/tiancai110a/gin-blog/pkg/gredis"
	"github.com/tiancai110a/gin-blog/service/cache_service"
)

type Article struct {
	Id            int
	TagId         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	CreatedBy     string
	ModifiedBy    string
	State         int

	PageNum  int
	PageSize int
}

func (a *Article) Get() (*models.Article, error) {
	//var cacheArticle *models.Article

	cache := cache_service.Article{Id: a.Id}
	key := cache.GetArticleKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			glog.Error("redis get wrong!", err)
			return nil, err
		}
		article := models.Article{}
		err = json.Unmarshal(data, &article)
		if err != nil {
			glog.Error("unmashal failed err:", err)
			return nil, err
		}
		return &article, nil
	}

	article, err := models.GetArticle(int(a.Id))
	if err != nil {
		glog.Error("db get article wrong! err:err", err)
		return nil, err
	}
	gredis.Set(key, article, 3600)

	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {

	cache := cache_service.Article{
		TagId:    a.TagId,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}

	key := cache.GetArticlesKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err == nil {
			var articles []*models.Article

			err = json.Unmarshal(data, articles)
			if err != nil {
				glog.Error("unmashal failed err:", err)
				return nil, err
			}
			return articles, nil

		}
		glog.Error("redis get wrong!", err)
	}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		glog.Error("redis get wrong!")
		return nil, err
	}
	gredis.Set(key, articles, 3600)
	return articles, nil
}
func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}
func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagId != -1 {
		maps["tag_id"] = a.TagId
	}

	return maps
}

func (a *Article) Add() error {

	article := models.Article{
		TagId:         a.TagId,
		Title:         a.Title,
		Desc:          a.Desc,
		Content:       a.Content,
		CreatedBy:     a.CreatedBy,
		State:         a.State,
		CoverImageUrl: a.CoverImageUrl,
	}

	if err := models.AddArticle(article); err != nil {
		glog.Error("AddArticle failed err:", err)
		return err
	}
	return nil
}

func (a *Article) ExistById() bool {
	return models.ExistArticleById(a.Id)
}

func (a *Article) Delete() error {
	cache := cache_service.Article{Id: a.Id}
	key := cache.GetArticlesKey()
	gredis.Delete(key)
	return models.DeleteArticle(a.Id)
}
