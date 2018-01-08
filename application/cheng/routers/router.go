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
	beego.Router("/insert", &controllers.Test{}, "post:Insert")
	beego.Router("/delete", &controllers.Test{}, "post:Delete")
	beego.Router("/update", &controllers.Test{}, "get:Update")
	beego.Router("/readLabel", &controllers.Test{}, "post:ReadLabel")
	beego.Router("deleteTest", &controllers.Test{}, "post:DeleteTest")
	beego.Router("/readTitleContent", &controllers.Test{}, "post:ReadTitleContent")
}
