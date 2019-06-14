package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, ok := strconv.ParseInt(c.Query("page"), 10, 64)
	if !ok {
		return 0
	}
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
