package routers

import (
	"github.com/tiancai110a/gin-blog/pkg/qrcode"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/tiancai110a/gin-blog/api"
	v1 "github.com/tiancai110a/gin-blog/api/v1"
	"github.com/tiancai110a/gin-blog/middleware/jwt"
	"github.com/tiancai110a/gin-blog/pkg/setting"
	"github.com/tiancai110a/gin-blog/pkg/upload"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/upload", api.UploadImage)

	r.GET("/auth", v1.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{

		//标签
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//文章
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		//二维码
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)

	}

	return r
}
