package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Init() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")
	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

//Translate 翻译错误信息
func Translate(err error) map[string][]string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}

type Employee1 struct {
	UserName string `json:"userName" binding:"required"`
	NickName string `json:"nickName" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=1,lte=120"`
	Email    string `json:"email" binding:"required,email"`
}

func main() {
	Init()

	r := gin.Default()
	r.POST("/zh", func(c *gin.Context) {
		var e Employee1
		err := c.ShouldBindJSON(&e)
		if err == nil {
			c.JSON(200, gin.H{"message": "Success"})
			return
		} else {
			c.JSON(200, gin.H{"message": Translate(err)})
			return
		}
	})
	r.Run()
}
