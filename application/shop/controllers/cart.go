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
 *     Initial: 2018/04/23        Shi Ruitao
 */

package controllers

import (
	"log"
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/shop/models"
	"github.com/shiruitao/go-one/application/shop/common"
)

type CareController struct {
	beego.Controller
}

func (this *CareController) AddCart() {
	var cart models.Cart

	userID := this.GetSession(common.SessionUserID).(uint32)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &cart)
	if err != nil {
		log.Println("You don't have access")
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSession}
	} else {
		cart.UserID = userID
		_, err = models.CartService.AddCart(&cart)
		if err != nil {
			log.Println("ErrMysql:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *CareController) ModifyNum() {
	var num struct {
		ID uint64 `json:"id"`
		Number int8 `json:"number"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &num)
	if err != nil {
		log.Println("You don't have access")
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSession}
	} else {
		_, err = models.CartService.ModifyNum(num.ID, num.Number)
		if err != nil {
			log.Println("ErrMysql:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *CareController) DeleteCart() {
	var id struct {
		ID uint64 `json:"id"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Println("You don't have access")
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSession}
	} else {
		_, err := models.CartService.DeleteWare(id.ID)
		if err != nil {
			log.Println("ErrMysql:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}
