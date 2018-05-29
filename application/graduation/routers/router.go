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

	// Admin
	beego.Router("/graduation/admin/addinfo", &controllers.UserController{}, "post:CreateUserInfo")
	beego.Router("/graduation/admin/modifyinfo", &controllers.UserController{}, "post:ModifyUserInfo")
	beego.Router("/graduation/admin/gettopic", &controllers.TopicController{}, "post:AdminGetTopic")
	beego.Router("/graduation/admin/check", &controllers.TopicController{}, "post:AdminCheck")

	// Teacher
	beego.Router("/graduation/teacher/add", &controllers.TopicController{}, "post:CreateTopic")
	beego.Router("/graduation/teacher/get", &controllers.TopicController{}, "post:TeacherGetTopic")
	beego.Router("/graduation/teacher/modify", &controllers.TopicController{}, "post:TeacherModifyTopic")
	beego.Router("/graduation/teacher/verify", &controllers.TopicController{}, "post:TeacherVerify")

	//Student
	beego.Router("/graduation/student/select", &controllers.TopicController{}, "post:Select")
	beego.Router("/graduation/student/back", &controllers.TopicController{}, "post:Back")
	beego.Router("/graduation/student/get", &controllers.TopicController{}, "post:StudentGetTopic")
}
