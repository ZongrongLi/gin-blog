package main

import (
	"testing"

	"github.com/tiancai110a/gin-blog/models"
)

func TestPreload(t *testing.T) {
	articles := []models.Article{}
	models.Testdb.Debug().Preload("Tag").Offset(0).Limit(10).Find(&articles)
}
