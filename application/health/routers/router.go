package routers

import (
	"github.com/shiruitao/go-one/application/health/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/health/user/login", &controllers.UserController{}, "post:CreateUser")

	beego.Router("/health/record/get", &controllers.RecordController{}, "post:Get")

	beego.Router("/health/user/test", &controllers.TestController{}, "post:Get")

	beego.Router("/health/question/add", &controllers.QuestionController{}, "post:Add")

}
