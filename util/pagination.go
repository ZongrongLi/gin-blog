package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tiancai110a/gin-blog/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		return 0
	}
	if page > 0 {
		result = int((page - 1) * int64(setting.PageSize))
	}

	return result
}
