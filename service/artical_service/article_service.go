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
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	CreatedBy     string
	ModifiedBy    string
	State         int
}

func (a *Article) Get() (*models.Article, error) {
	//var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.Id}
	key := cache.GetArticleKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			glog.Error("redis get wrong!")
			return nil, err
		}
		article := models.Article{}
		json.Unmarshal(data, &article)
		return &article, nil
	}

	article, err := models.GetArticle(int(a.Id))
	if err != nil {
		glog.Error("redis get wrong!")
		return nil, err
	}
	gredis.Set(key, article, 3600)

	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {

}
