// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/shiruitao/go-one/application/cheng/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.Test{}, "get:Helloworld")
	beego.Router("/add", &controllers.Test{}, "get:Add")
	beego.Router("/insert", &controllers.Test{}, "get:Insert")
	beego.Router("/read", &controllers.Test{}, "get:Read")
}
