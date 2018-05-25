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

	"github.com/shiruitao/go-one/application/graduation/controllers"
)

func init() {
	beego.Router("/graduation/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/graduation/user/changepassword", &controllers.UserController{}, "post:ChangePassword")
	beego.Router("/graduation/user/addinfo", &controllers.UserController{}, "post:CreateUserInfo")
	beego.Router("/graduation/user/modifyinfo", &controllers.UserController{}, "post:ModifyUserInfo")
}
