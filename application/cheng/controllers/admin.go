/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Inc.
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

package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/shiruitao/go-one/application/cheng/common"
	"github.com/shiruitao/go-one/application/cheng/log"
	"github.com/shiruitao/go-one/application/cheng/models"
)

type AdminController struct {
	beego.Controller
}

// 管理注册
func (this *AdminController) Create() {
	var admin models.Admin
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
	if err != nil {
		log.Logger.Error("json.Unmarshal", err)
	} else {
		err := models.AdminService.Create(admin)
		if err != nil {
			log.Logger.Error("models.Insert", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

// 管理登录
func (this *AdminController) Login() {
	var admin models.Admin
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		flag, err := models.AdminService.Login(admin.Name, admin.Password)
		if err != nil {
			if err == orm.ErrNoRows {
				log.Logger.Warn("Invalid name!")
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidUser}
			} else {
				log.Logger.Error("error", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			if !flag {
				log.Logger.Debug("Wrong Password!")
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrWrongPass}
			} else {
				this.SetSession(common.SessionAdminID, admin.Name)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	}
	this.ServeJSON()
}

// 根据name修改管理密码
func (this *AdminController) ChangePass() {
	var admin struct {
		Oldpass string
		Newpass string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		name := this.GetSession(common.SessionAdminID)
		flag, err := models.AdminService.Login(name.(string), admin.Oldpass)

		if err != nil {
			log.Logger.Error("Old Password:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			if !flag {
				log.Logger.Debug("Wrong Password!")
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrWrongPass}
			} else {
				err := models.AdminService.ChangePass(name.(string), admin.Newpass)

				if err != nil {
					log.Logger.Error("models.ChangePass:", err)

					this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
				} else {
					this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
				}
			}
		}
	}
	this.ServeJSON()
}
