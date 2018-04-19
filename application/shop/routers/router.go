package routers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/shop/controllers"
)

func init() {
    beego.Router("/shop/user/create", &controllers.UserController{}, "post:CreateUser")

	beego.Router("/shop/ware/create", &controllers.WareController{}, "post:CreateWare")
}
