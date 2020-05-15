package response

const (
	SUCCESS                  = 200
	ERROR                    = 500
	INVALID_PARAMS           = 400
	NOFIND_FILE_OR_DIRECTORY = 20501
	ERROR_PERMISSION         = 20502

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_CREATE_EXIST_USERNAME    = 20002
	ERROR_CREATE_EXIST_TELEPHONE   = 20003
	ERROR_CREATE_NAME_OR_PASSWORD  = 20004
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 200022
	ERROR_AUTH_TOKEN               = 20010
	ERROR_AUTH                     = 20011
	ERROR_EMPTY_AUTH               = 20012
)

var MsgFlags = map[int]string{
	SUCCESS:                  "ok",
	ERROR:                    "fail",
	INVALID_PARAMS:           "请求参数错误",
	NOFIND_FILE_OR_DIRECTORY: "文件夹或者见目录不存在",
	ERROR_PERMISSION:         "没有权限",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_CREATE_EXIST_USERNAME:    "用户名已存在",
	ERROR_CREATE_EXIST_TELEPHONE:   "此手机号已注册",
	ERROR_CREATE_NAME_OR_PASSWORD:  "用户名或密码错误",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_EMPTY_AUTH:               "Token为空",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
