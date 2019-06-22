package cache_service

import (
	"strconv"
	"strings"
)

type Article struct {
	Id    int
	TagId int
	State int

	PageNum  int
	PageSize int
}

const (
	CACHE_ARTICLE = "ARTICLE"
	CACHE_TAG     = "TAG"
)

func (a *Article) GetArticleKey() string {
	return CACHE_ARTICLE + "_" + strconv.Itoa(a.Id)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		CACHE_ARTICLE,
		"LIST",
	}

	if a.Id > 0 {
		keys = append(keys, strconv.Itoa(a.Id))
	}
	if a.TagId > 0 {
		keys = append(keys, strconv.Itoa(a.TagId))
	}
	if a.State >= 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
