package errorx

import (
	"ddd-template/common/conf"
	_ "embed"
)

//go:embed code.go
var ByteCodeFile []byte

type CodeType = int

const (
	ServerError        CodeType = 10101
	TooManyRequests    CodeType = 10102
	ParamBindError     CodeType = 10103
	AuthorizationError CodeType = 10104
	UrlSignError       CodeType = 10105
	CacheSetError      CodeType = 10106
	CacheGetError      CodeType = 10107
	CacheDelError      CodeType = 10108
	CacheNotExist      CodeType = 10109
	ResubmitError      CodeType = 10110
	HashIdsEncodeError CodeType = 10111
	HashIdsDecodeError CodeType = 10112
	RBACError          CodeType = 10113
	RedisConnectError  CodeType = 10114
	MySQLConnectError  CodeType = 10115
	WriteConfigError   CodeType = 10116
	SendEmailError     CodeType = 10117
	MySQLExecError     CodeType = 10118
	GoVersionError     CodeType = 10119
	SocketConnectError CodeType = 10120
	SocketSendError    CodeType = 10121

	AuthorizedCreateError    CodeType = 20101
	AuthorizedListError      CodeType = 20102
	AuthorizedDeleteError    CodeType = 20103
	AuthorizedUpdateError    CodeType = 20104
	AuthorizedDetailError    CodeType = 20105
	AuthorizedCreateAPIError CodeType = 20106
	AuthorizedListAPIError   CodeType = 20107
	AuthorizedDeleteAPIError CodeType = 20108

	AdminCreateError             CodeType = 20201
	AdminListError               CodeType = 20202
	AdminDeleteError             CodeType = 20203
	AdminUpdateError             CodeType = 20204
	AdminResetPasswordError      CodeType = 20205
	AdminLoginError              CodeType = 20206
	AdminLogOutError             CodeType = 20207
	AdminModifyPasswordError     CodeType = 20208
	AdminModifyPersonalInfoError CodeType = 20209
	AdminMenuListError           CodeType = 20210
	AdminMenuCreateError         CodeType = 20211
	AdminOfflineError            CodeType = 20212
	AdminDetailError             CodeType = 20213

	MenuCreateError       CodeType = 20301
	MenuUpdateError       CodeType = 20302
	MenuListError         CodeType = 20303
	MenuDeleteError       CodeType = 20304
	MenuDetailError       CodeType = 20305
	MenuCreateActionError CodeType = 20306
	MenuListActionError   CodeType = 20307
	MenuDeleteActionError CodeType = 20308

	CronCreateError  CodeType = 20401
	CronUpdateError  CodeType = 20402
	CronListError    CodeType = 20403
	CronDetailError  CodeType = 20404
	CronExecuteError CodeType = 20405
)

func Text(code CodeType) string {
	lang := conf.Get().Language

	if lang == "zh-cn" {
		return zhCNText[code]
	}

	if lang == "en-us" {
		return enUSText[code]
	}

	return zhCNText[code]
}
