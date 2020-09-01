package errmsg

const (
	//通用业务错误
	YES = 10
	NO = 50
	Success = 200
	ERROR   = 500

	//用户模块错误码 code = 1000...
	ERROR_USERNAME_USED   = 1001
	ERROR_PASSWORD_WRONG  = 1002
	ERROR_USER_NOT_EXIST  = 1003
	ERROR_TOKEN_NOT_EXIST = 1004
	ERROR_TOKEN_TIMEOUT   = 1005
	ERROR_TOKEN_WRONG     = 1006
	ERROR_USER_NOT_AUTHORIZED = 1007
	//文章模块错误 code = 2000..。
	ERROR_ARTICE_NOT_EXIST = 2001
	//分类模块错误 code = 3000...
	ERROR_CATEGORYNAME_USED = 3001
	ERROR_CATEGORY_NOT_EXIST = 3002
)

var codemsg = map[int]string{
	Success: "ok",
	ERROR:   "fail",

	ERROR_USERNAME_USED:   "用户名已被使用",
	ERROR_PASSWORD_WRONG:  "密码错误",
	ERROR_USER_NOT_EXIST:  "用户不存在",
	ERROR_USER_NOT_AUTHORIZED: "用户权限不足",

	ERROR_TOKEN_NOT_EXIST: "token不存在",
	ERROR_TOKEN_TIMEOUT:   "toke超时",
	ERROR_TOKEN_WRONG:     "token错误",

	ERROR_CATEGORY_NOT_EXIST: "类别不存在",
	ERROR_CATEGORYNAME_USED: "类别名已被使用",

	ERROR_ARTICE_NOT_EXIST: "文章不存在",
}

func Errmsg(code int) string {
	return codemsg[code]
}