// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/community/controllers"
)

func init() {
	beego.Router("/community/user/create", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/community/user/modify", &controllers.UserController{}, "post:UserUpdate")
	beego.Router("/community/user/getuser", &controllers.UserController{}, "get:GetUser")
	beego.Router("/community/user/getusername", &controllers.UserController{}, "post:GetUserByName")
	beego.Router("/community/user/getuserarea", &controllers.UserController{}, "post:GetUserByArea")
	beego.Router("/community/user/getuserage", &controllers.UserController{}, "post:GetUserByAge")
}
