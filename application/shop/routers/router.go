package routers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/shop/controllers"
)

func init() {
	// User
    beego.Router("/shop/user/create", &controllers.UserController{}, "post:CreateUser")

    // Ware
	beego.Router("/shop/ware/create", &controllers.WareController{}, "post:CreateWare")
	beego.Router("/shop/ware/update", &controllers.WareController{}, "post:UpdateWare")
	beego.Router("/shop/ware/getall", &controllers.WareController{}, "get:GetAll")
	beego.Router("/shop/ware/recommend", &controllers.WareController{}, "get:GetRecommend")
	beego.Router("/shop/ware/status", &controllers.WareController{}, "post:StatusWare")

	// Cart

	// Order

	// Address
	beego.Router("/shop/address/add", &controllers.AddressController{}, "post:AddAddress")
}
