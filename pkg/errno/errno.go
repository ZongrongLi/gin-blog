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
)

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}
