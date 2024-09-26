package validate

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"siteol.com/smart/src/common/log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

// TransLangSupport 支持更多语言请添加
var TransLangSupport = []string{"zh-CN", "en-US"}
var v *validator.Validate                        // 自定义校验器使用
var transMap = make(map[string]ut.Translator, 0) // 校验错误取这个Map进行错误翻译

// init
func init() {
	validate, err := transInit()
	if err != nil {
		log.ErrorF("transInitValidate Fail")
		return
	}
	v = validate
}

// 初始化语言国际化
func transInit() (validate *validator.Validate, err error) {
	if vObj, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(zhT, zhT, enT)
		validate = vObj
		// 注册自定义校验
		CustomBindingValidator(validate)
		// 初始化全部翻译对象
		for _, local := range TransLangSupport {
			initLocal := local[:strings.Index(local, "-")]
			trans, o := uni.GetTranslator(initLocal)
			if !o {
				err = errors.New("GetTranslator Err")
				return
			}
			// 注册全部翻译器
			switch initLocal {
			case "en":
				err = enTrans.RegisterDefaultTranslations(validate, trans)
			//case "zh":
			//	err = zhTrans.RegisterDefaultTranslations(validate, translator)
			default:
				err = zhTrans.RegisterDefaultTranslations(validate, trans)
			}
			// 注册自定义文言
			CustomBindingTranslations(validate, &trans)
			transMap[local] = trans
		}
	}
	return
}

// CustomBindingValidator 注册自定义校验
func CustomBindingValidator(validate *validator.Validate) {
	registryValidator(validate, "letterUnder", letterUnder)
}

// CustomBindingTranslations 注册翻译器
func CustomBindingTranslations(validate *validator.Validate, trans *ut.Translator) {
	validate.RegisterTranslation("letterUnder", *trans, returnRegTransFunc("letterUnder"), returnTransFunc("letterUnder"))
}

func registryValidator(validate *validator.Validate, name string, fn validator.Func) {
	err := validate.RegisterValidation(name, fn)
	if err != nil {
		return
	}
}

func returnRegTransFunc(name string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		return ut.Add(name, returnTrans(name, ut.Locale()), true)
	}
}

func returnTransFunc(name string) validator.TranslationFunc {
	return func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(name, fe.Field())
		return t
	}
}
