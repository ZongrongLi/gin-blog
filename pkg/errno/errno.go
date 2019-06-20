package errno

var (
	// Common errors
	Success       = &Errno{Code: 0, Message: "SUCCESS"}
	Error         = &Errno{Code: 10001, Message: "ERROR"}
	InvalidParams = &Errno{Code: 10002, Message: "INVALID_PARAMS"}
	ErrorExistTag = &Errno{Code: 10005, Message: "ERROR_EXIST_TAG."}

	// user errors
	ErrorNotexistTag           = &Errno{Code: 20103, Message: "ERROR_NOT_EXIST_TAG"}
	ErrorNotexistArticle       = &Errno{Code: 20104, Message: "ERROR_NOT_EXIST_ARTICLE."}
	ErrorAuthCheckTokenFail    = &Errno{Code: 20106, Message: "ERROR_AUTH_CHECK_TOKEN_FAIL."}
	ErrorAuthCheckTokenTimeOut = &Errno{Code: 20107, Message: "ERROR_AUTH_CHECK_TOKEN_TIMEOUT."}
	ErrorAuthToken             = &Errno{Code: 20102, Message: "ERROR_AUTH_TOKEN."}
	ErrorAuth                  = &Errno{Code: 20105, Message: "ERROR_AUTH"}

	//upload
	// 检查图片失败
	ErrorUploadCheckImageFail = &Errno{Code: 20106, Message: "ERROR_UPLOAD_CHECK_IMAGE_FAIL"}
	// 校验图片错误，图片格式或大小有问题
	ErrorUploadCheckImageFormat = &Errno{Code: 20107, Message: "ERROR_UPLOAD_CHECK_IMAGE_FORMAT"}
	// 保存图片失败
	ErrorUploadSaveImageFail = &Errno{Code: 20108, Message: "ERROR_UPLOAD_SAVE_IMAGE_FAIL"}
)

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}
