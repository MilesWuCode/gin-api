package plugin

import (
	en "github.com/go-playground/locales/en"
	zh_tW "github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	zh_tW_translations "github.com/go-playground/validator/v10/translations/zh_tw"
	"reflect"
	"strings"
)

func Validate(s interface{}) validator.ValidationErrorsTranslations {
	zh_tW := zh_tW.New()

	en := en.New()

	uni := ut.New(zh_tW, en)

	trans, _ := uni.GetTranslator("zh_tw")

	// trans.Add("Name", "名稱", false)
	// trans.Add("email", "帳號", false)

	validate := validator.New()

	// validate.RegisterTagNameFunc(func(field reflect.StructField) string {
	// 	label := field.Tag.Get("label")
	// 	if label == "" {
	// 		return field.Name
	// 	}
	// 	return label
	// })

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	zh_tW_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(s)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		return errs.Translate(trans)
	} else {
		return nil
	}
}
