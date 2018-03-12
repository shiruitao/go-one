package routers

import (
	"github.com/shiruitao/go-one/application/shop/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
