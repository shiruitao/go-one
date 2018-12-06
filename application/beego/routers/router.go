// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/shiruitao/go-one/application/beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),

		beego.NSNamespace("/document",
			beego.NSInclude(
				&controllers.DocumentController{},
			),
		),

		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),

		beego.NSNamespace("/recommend",
			beego.NSInclude(
				&controllers.RecommendController{},
			),
		),

		beego.NSNamespace("/slideshow",
			beego.NSInclude(
				&controllers.SlideshowController{},
			),
		),

		beego.NSNamespace("/teacher",
			beego.NSInclude(
				&controllers.TeacherController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
