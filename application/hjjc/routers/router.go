package routers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/hjjc/controllers"
)

func init() {
	beego.Router("/hjjc/user/add", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/hjjc/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/hjjc/user/get", &controllers.UserController{}, "post:Get")
	beego.Router("/hjjc/user/update", &controllers.UserController{}, "post:Update")
	beego.Router("/hjjc/user/delete", &controllers.UserController{}, "post:Delete")

	beego.Router("/hjjc/device/get", &controllers.DeviceController{}, "post:Get")

	beego.Router("/hjjc/devices/get", &controllers.DevicesController{}, "post:Get")
	beego.Router("/hjjc/devices/add", &controllers.DevicesController{}, "post:Add")
	beego.Router("/hjjc/devices/update", &controllers.DevicesController{}, "post:Update")
	beego.Router("/hjjc/devices/delete", &controllers.DevicesController{}, "post:Delete")

	beego.Router("/hjjc/company/get", &controllers.CompanyController{}, "post:Get")
	beego.Router("/hjjc/company/add", &controllers.CompanyController{}, "post:Add")
	beego.Router("/hjjc/company/update", &controllers.CompanyController{}, "post:Update")
	beego.Router("/hjjc/company/delete", &controllers.CompanyController{}, "post:Delete")

	beego.Router("/hjjc/abnormal/get", &controllers.AbnormalController{}, "post:Get")

	beego.Router("/hjjc/analyses/get", &controllers.AnalysesController{}, "post:Average")

}
