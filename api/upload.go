package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/tiancai110a/gin-blog/pkg/errno"
	"github.com/tiancai110a/gin-blog/pkg/upload"
)

func UploadImage(c *gin.Context) {
	errnumber := errno.Success
	data := make(map[string]string)

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": errnumber.Code,
			"msg":  errnumber.Message,
			"data": data,
		})
	}()

	file, image, err := c.Request.FormFile("image")
	if err != nil || image == nil {
		errnumber = errno.InvalidParams
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	src := fullPath + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		errnumber = errno.ErrorUploadCheckImageFormat
		return
	}
	err = upload.CheckImage(fullPath)
	if err != nil {
		glog.Warning(err)
		errnumber = errno.ErrorUploadCheckImageFail
		return
	}
	if err := c.SaveUploadedFile(image, src); err != nil {
		glog.Warning(err)
		errnumber = errno.ErrorUploadSaveImageFail
		return
	}
	data["image_url"] = upload.GetImageFullUrl(imageName)
	data["image_save_url"] = savePath + imageName
	return
}
