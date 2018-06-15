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
	beego.Router("/bbs/user/register", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/bbs/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/bbs/user/changepassword", &controllers.UserController{}, "post:ChangePassword")
	beego.Router("/bbs/user/get", &controllers.UserController{}, "get:Get")
	beego.Router("/bbs/user/delete", &controllers.UserController{}, "post:Delete")

	beego.Router("/bbs/art/add", &controllers.ArticleController{}, "post:ArticleCreate")
	beego.Router("/bbs/art/get", &controllers.ArticleController{}, "get:Get")
	beego.Router("/bbs/art/deleteart", &controllers.ArticleController{}, "post:DeleteArt")
	beego.Router("/bbs/art/deleteuser", &controllers.ArticleController{}, "post:DeleteUser")

	beego.Router("/bbs/comment/addrep", &controllers.CommentController{}, "post:AddRep")
	beego.Router("/bbs/comment/addcreator", &controllers.CommentController{}, "post:AddCreator")
	beego.Router("/bbs/comment/get", &controllers.CommentController{}, "post:Get")
}
