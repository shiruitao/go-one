package routers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/hjjc/controllers"
)

func init() {
	beego.Router("/hjjc/device/get", &controllers.DeviceController{}, "post:Get")

	beego.Router("/hjjc/company/get", &controllers.CompanyController{}, "post:Get")
}