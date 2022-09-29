package schema

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

type Error struct {
	Failed string      `json:"failed"`
	Tag    string      `json:"tag"`
	Value  interface{} `json:"value"`
	Detail string      `json:"detail"`
}

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()

}

func translate(language string) ut.Translator {
	var trans ut.Translator
	switch language {
	case "en":
		defaultEn := en.New()
		uni := ut.New(defaultEn, defaultEn)
		trans, _ = uni.GetTranslator(language)
		_ = en_translation.RegisterDefaultTranslations(validate, trans)
	default:
		defaultZh := zh.New()
		uni := ut.New(defaultZh, defaultZh)
		trans, _ = uni.GetTranslator(language)
		_ = zh_translation.RegisterDefaultTranslations(validate, trans)
	}
	return trans
}

func ValidateForm(c interface{}, language string) map[string]*Error {
	var (
		errs = map[string]*Error{}
	)
	err := validate.Struct(c)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			_err := &Error{}
			_err.Detail = e.Translate(translate(language))
			_err.Tag = e.Tag()
			_err.Failed = e.Field()
			_err.Value = e.Value()
			errs[strings.ToLower(e.Field()[:1])+e.Field()[1:]] = _err
		}
	}
	return errs
}
