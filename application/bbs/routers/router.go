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

	"github.com/shiruitao/go-one/application/bbs/controllers"
)

func init() {
	beego.Router("/bbs/user/register", &controllers.UserController{}, "post:RegisterUser")
	beego.Router("/bbs/user/login", &controllers.UserController{}, "post:LoginUser")
	beego.Router("/bbs/user/changepassword", &controllers.UserController{}, "post:ChangePassword")
	beego.Router("/bbs/user/changeinfo", &controllers.UserController{}, "post:ChangeInfo")
}
