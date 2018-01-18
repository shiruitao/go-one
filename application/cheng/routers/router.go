/*
 * MIT License
 *
 * Copyright (c) 2018 Shi Ruitao.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/01/02        Shi Ruitao
 */

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
	"github.com/shiruitao/go-one/application/cheng/controllers"
)

func init() {
	// 全读
	beego.Router("/readAll", &controllers.Test{}, "get:ReadAll")
	// 根据标签查询
	beego.Router("/readLabel", &controllers.Test{}, "post:ReadLabel")
	// 根据标题和内容查询
	beego.Router("/readTitleContent", &controllers.Test{}, "post:ReadTitleContent")
	// 根据时间查询
	beego.Router("/readTime", &controllers.Test{}, "post:ReadTime")

	// 发表文章
	beego.Router("/insert", &controllers.Test{}, "post:Insert")
	// 删除文章
	beego.Router("/delete", &controllers.Test{}, "post:Delete")
	// 注册
	beego.Router("/adminCreate", &controllers.AdminController{}, "post:Create")
	// 登录
	beego.Router("/adminLogin", &controllers.AdminController{}, "post:Login")
	// 更改密码
	beego.Router("/adminChangePass", &controllers.AdminController{}, "post:ChangePass")

	beego.Router("/deleteTest", &controllers.Test{}, "post:DeleteTest")
}
