package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Student struct {
	Name       string    `form:"name" binding:"required"`
	Score      int       `form:"score" binding:"gt=0"`
	Enrollment time.Time `form:"enrollment" binding:"required,before_today" time_format:"2006-01-02" time_utc:"8"`
	Graduation time.Time `form:"graduation" binding:"required,gtfield=Enrollment" time_format:"2006-01-02" time_utc:"8"`
}

// 自定义验证器
var beforeToday validator.Func = func(f1 validator.FieldLevel) bool {
	if date, ok := f1.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Before(today) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func processErr(err error) {
	if err == nil {
		return
	}

	//给Validate.Struct()函数传了一个非法的参数
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return
	}

	validationErrs := err.(validator.ValidationErrors)
	for _, validationErr := range validationErrs {
		fmt.Printf("field %s 不满足条件 %s\n", validationErr.Field(), validationErr.Tag())
	}
}

func main() {

	engine := gin.Default()

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("before_today", beforeToday)
	}

	engine.GET("/", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBind(&stu); err != nil {
			processErr(err)
			ctx.String(http.StatusBadRequest, "parse parameter failed")
		} else {
			ctx.JSON(http.StatusOK, stu)
		}
	})

	engine.Run(":5656")
	//{{base_url}}?name=zcy&score=1&enrollment=2021-08-23&graduation=2021-09-23
}
