package tag_service

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/tiancai110a/gin-blog/models"
	"github.com/tiancai110a/gin-blog/pkg/gredis"
	"github.com/tiancai110a/gin-blog/service/cache_service"
)

type Tag struct {
	Id         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if t.State != -1 {
		maps["state"] = t.State
	}
	if t.Id != -1 {
		maps["tag_id"] = t.Id
	}

	return maps
}

func (t *Tag) ExistByName() bool {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistById() bool {
	return models.ExistTagById(t.Id)
}

func (t *Tag) GetAll() ([]models.Tag, error) {

	tag := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := tag.GetTagsKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err == nil {
			var tags []models.Tag
			err = json.Unmarshal(data, &tags)
			if err != nil {
				glog.Error("redis tags data unmarshal failed err:", err)
			} else {
				return tags, nil
			}
		}

		glog.Error("redis get tags failed err:", err)
	}

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())

	if err != nil {
		glog.Error("db get tags failed err:", err)
		return nil, err
	}
	gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) Get() (*models.Tag, error) {
	cache := cache_service.Tag{ID: t.Id}
	key := cache.GetTagsKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err == nil {
			tag := models.Tag{}
			err := json.Unmarshal(data, &tag)
			if err != nil {
				glog.Error("redis data unmashal failed err:", err)
			} else {
				return &tag, nil
			}
		}
		glog.Error("redis get wrong!", err)
	}

	tag, err := models.GetTag(int(t.Id))
	if err != nil {
		glog.Error("db get Tag wrong! err:err", err)
		return nil, err
	}
	gredis.Set(key, tag, 3600)
	return tag, nil
}

func (t *Tag) Count() (int, error) {
	return models.GetArticleTotal(t.getMaps())
}

func (t *Tag) Add() error {

	tag := models.Tag{
		Name:      t.Name,
		CreatedBy: t.CreatedBy,
		State:     t.State,
	}

	if err := models.AddTag(&tag); err != nil {
		glog.Error("AddArticle failed err:", err)
		return err
	}

	cache := cache_service.Tag{ID: t.Id}
	key := cache.GetTagsKey()
	gredis.Set(key, tag, 3600)
	return nil
}

func (t *Tag) Edit() error {

	tag := models.Tag{
		Name:       t.Name,
		ModifiedBy: t.ModifiedBy,
		State:      t.State,
	}

	if err := models.EditTag(t.Id, tag); err != nil {
		glog.Error("AddArticle failed err:", err)
		return err
	}

	cache := cache_service.Tag{ID: t.Id}
	key := cache.GetTagsKey()
	gredis.Set(key, tag, 3600)
	return nil
}

func (t *Tag) Delete() {
	cache := cache_service.Tag{ID: t.Id}
	key := cache.GetTagsKey()
	gredis.Delete(key)

}
