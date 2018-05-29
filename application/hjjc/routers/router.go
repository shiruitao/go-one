package routers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/hjjc/controllers"
)

func init() {
	beego.Router("/community/user/create", &controllers.UserController{}, "post:CreateUser")

	beego.Router("/hjjc/device/get", &controllers.DeviceController{}, "post:Get")
}
