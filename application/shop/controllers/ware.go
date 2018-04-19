/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
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
 *     Initial: 2018/04/19        Shi Ruitao
 */

package controllers

import (
	"log"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"

	"github.com/shiruitao/go-one/application/shop/common"
	"github.com/shiruitao/go-one/application/shop/models"
	"github.com/shiruitao/go-one/application/shop/util"
)

type WareController struct {
	beego.Controller
}

func (this *WareController) CreateWare() {
	var ware models.Ware

	isAdmin := this.Ctx.Request.Context().Value("user").(jwt.MapClaims)[util.IsAdmin].(bool)
	if !isAdmin {
		log.Println("You don't have access")
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrToken}
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ware)
	if err != nil {
		log.Println("error json:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	}

	id, err := models.WareService.AddWare(&ware)
	if err != nil {
		log.Println("error mysql:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "id": id}
	}
	this.ServeJSON()
}
