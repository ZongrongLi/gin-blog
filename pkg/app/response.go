package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tiancai110a/gin-blog/pkg/errno"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, errCode *errno.Errno, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode.Code,
		"msg":  errCode.Message,
		"data": data,
	})

	return
}
