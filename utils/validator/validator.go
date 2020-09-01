package validator

import (
	"blog/utils/errmsg"
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unitrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

//传入data的数据需要是结构体
func Validate(data interface{}) (string,int) {
	validate := validator.New()
	utrans := unitrans.New(zh_Hans_CN.New())
	tran, _ := utrans.GetTranslator("zh_Hans_cn")

	err := zh.RegisterDefaultTranslations(validate, tran)
	if err != nil {
		fmt.Println(err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(tran), errmsg.ERROR
		}
	}
	return "",errmsg.Success
}
