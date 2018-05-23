package routers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/shop/controllers"
)

func init() {
	// User
	beego.Router("/shop/user/login", &controllers.UserController{}, "post:CreateUser")

	// Ware
	beego.Router("/shop/ware/create", &controllers.WareController{}, "post:CreateWare")
	beego.Router("/shop/ware/update", &controllers.WareController{}, "post:UpdateWare")
	beego.Router("/shop/ware/getall", &controllers.WareController{}, "get:GetAll")
	beego.Router("/shop/ware/recommend", &controllers.WareController{}, "get:GetRecommend")
	beego.Router("/shop/ware/status", &controllers.WareController{}, "post:StatusWare")

	// Cart
	beego.Router("/shop/cart/add", &controllers.CartController{}, "post:AddCart")
	beego.Router("/shop/cart/number", &controllers.CartController{}, "post:ModifyNum")
	beego.Router("/shop/cart/delete", &controllers.CartController{}, "post:DeleteCart")

	// Order
	beego.Router("/shop/order/add", &controllers.OrderController{}, "post:AddOrder")
	beego.Router("/shop/order/finish", &controllers.OrderController{}, "post:FinishOrder")
	beego.Router("/shop/order/get", &controllers.OrderController{}, "post:GetOrder")

	// Address
	beego.Router("/shop/address/add", &controllers.AddressController{}, "post:AddAddress")
}
