package internal

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	unit "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
)

const (
	ZH Locale = "zh"
	EN Locale = "en"
)

type Locale string

const (
	requiredTrimTag = "required_trim"
	onlyTrimTag     = "trim"
)

var (
	trans unit.Translator
)

func InitTrans(locale Locale) error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil
	}

	if err := v.RegisterValidation(requiredTrimTag, requiredTrim); err != nil {
		return err
	}

	if err := v.RegisterValidation(onlyTrimTag, onlyTrimSpace); err != nil {
		return err
	}

	zhT := zh.New()
	enT := en.New()

	uni := unit.New(enT, zhT, enT)

	trans, ok = uni.GetTranslator(string(locale))
	if !ok {
		return errors.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	var err error
	// 注册翻译器
	switch locale {
	case ZH:
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}

	if err := v.RegisterTranslation(requiredTrimTag, trans,
		registerTranslator(requiredTrimTag, "{0}为必填字段"),
		translationFunc); err != nil {
		return err
	}

	return err

}

func Translate(errs validator.ValidationErrors) map[string]string {
	return removeTopStruct(errs.Translate(trans))
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// requiredTrim 去空格后必填
func requiredTrim(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		v := strings.TrimSpace(field.String())
		if len(v) == 0 {
			return false
		}

		if fl.Parent().Kind() == reflect.Ptr {
			field = fl.Parent().Elem().FieldByName(fl.StructFieldName())
		} else {
			field = fl.Parent().FieldByName(fl.StructFieldName())
		}

		field.SetString(v)

		return true
	default:
		return true
	}

}

// onlyTrimSpace 仅去空格
func onlyTrimSpace(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		v := strings.TrimSpace(field.String())
		if fl.Parent().Kind() == reflect.Ptr {
			field = fl.Parent().Elem().FieldByName(fl.StructFieldName())
		} else {
			field = fl.Parent().FieldByName(fl.StructFieldName())
		}

		field.SetString(v)

	}

	return true
}

func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans unit.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

func translationFunc(trans unit.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
