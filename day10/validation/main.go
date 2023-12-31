package main

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var val = validator.New()

type RegistRequest struct {
	UserName   string `validate:"gt=0"`             //>0 长度大于0
	PassWord   string `validate:"min=6,max=12"`     //密码长度[6,12]
	PassRepeat string `validate:"eqfield=PassWord"` //跨字段相等校验
	Email      string `validate:"email"`            //需要满足email的格式
}

func validateEmail(f1 validator.FieldLevel) bool {

	input := f1.Field().String()
	if pass, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, input); pass {
		return true
	}
	return false

}

type InnerRequest struct {
	Pass  string `validate:"min=6,max=12"` //密码长度[6,12]
	Email string `validate:"school"`
}

type OutterRequest struct {
	PassWord   string `validate:"eqcsfield=Nest.Pass"` //跨结构体相等校验
	PassRepeat string `validate:"eqfield=PassWord"`    //跨字段相等校验
	Nest       InnerRequest
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

	//ValidationErrors是一个错误切片,它保存了每个字段违反的每个约束信息
	validationErrs := err.(validator.ValidationErrors)
	for _, validationErr := range validationErrs {
		fmt.Printf("field %s 不满足条件 %s\n", validationErr.Field(), validationErr.Tag())
	}

}

func main() {
	val.RegisterValidation("school", validateEmail) //注册一个自定义的validator

	req := RegistRequest{
		UserName:   "zcy",
		PassWord:   "123456",
		PassRepeat: "123456",
		Email:      "123qq.com",
	}
	processErr(val.Struct(req))
	processErr(val.Struct(3))

	fmt.Println("============")

	inreq := InnerRequest{
		Pass:  "1234567",
		Email: "123qq.com",
	}
	outreq := OutterRequest{
		PassWord:   "123456",
		PassRepeat: "1234568",
		Nest:       inreq,
	}
	processErr(val.Struct(outreq))

}
