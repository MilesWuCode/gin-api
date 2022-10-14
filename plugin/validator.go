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

func Validate(s interface{}) map[string]string {
	zh_tW := zh_tW.New()

	en := en.New()

	uni := ut.New(zh_tW, en)

	trans, _ := uni.GetTranslator("zh_tw")

	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		label := fld.Tag.Get("label")

		if label != "" {
			return label
		}

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
		errMsg := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.StructField()

			field, _ := reflect.TypeOf(s).FieldByName(fieldName)

			jsonKey := field.Tag.Get("json")

			errMsg[jsonKey] = err.Translate(trans)
		}

		return errMsg
	} else {
		return nil
	}
}
