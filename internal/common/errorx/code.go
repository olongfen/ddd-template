package errorx

import (
	"ddd-template/internal/common/conf"
	_ "embed"
)

type BizCode int

func Text(code BizCode) string {
	lang := conf.Get().Language

	if lang == "zh-cn" {
		return zhCNText[code]
	}

	if lang == "en-us" {
		return enUSText[code]
	}

	return zhCNText[code]
}

const (
	// IllegalAccessToken 非法token
	IllegalAccessToken BizCode = 40001
	// IllegalCertificate 非法凭证
	IllegalCertificate BizCode = 40002
	// IllegalParameter 非法参数
	IllegalParameter BizCode = 40003
)
