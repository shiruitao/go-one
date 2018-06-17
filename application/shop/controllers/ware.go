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
	"encoding/json"
	"log"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/shop/common"
	"github.com/shiruitao/go-one/application/shop/models"
)

type WareController struct {
	beego.Controller
}

func (this *WareController) CreateWare() {
	var ware models.Commodity

	//isAdmin := this.GetSession(common.SessionIsAdmin)
	//if isAdmin != true {
	//	log.Println("You don't have access")
	//	this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSession}
	//} else {
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ware)
	if err != nil {
		log.Println("error json:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		_, err = models.WareService.AddWare(&ware)
		if err != nil {
			log.Println("ErrMysql:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	//}
	this.ServeJSON()
}

func (this *WareController) GetAll() {
	ware, num, err := models.WareService.GetAll()
	if err != nil {
		log.Println("ErrMysql:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: ware, "number": num}
	}
	this.ServeJSON()
}

func (this *WareController) Delete() {
	var id struct {
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Println("error json:", err)
	}
	err = models.WareService.Delete(id.ID)
	if err != nil {
		log.Println("ErrMysql:", err)
	}
}

func (this *WareController) GetRecommend() {
	ware, num, err := models.WareService.GetRecommend()
	if err != nil {
		log.Println("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "num": num, common.RespKeyData: ware}
	}
	this.ServeJSON()
}

func (this *WareController) UpdateWare() {
	var (
		ware models.Commodity
	)

	isAdmin := this.GetSession(common.SessionIsAdmin)
	if isAdmin != true {
		log.Println("You don't have access")
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSession}
	} else {
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &ware)
		if err != nil {
			log.Println("error json:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		} else {
			err = models.WareService.UpdateWare(&ware)
			if err != nil {
				log.Println("ErrMysql:", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}

	}
	this.ServeJSON()
}

func (this *WareController) StatusWare() {
	var status struct {
		ID     uint32 `json:"id"`
		Status int8   `json:"status"`
	}

	isAdmin := this.GetSession(common.SessionIsAdmin)
	if isAdmin != true {
		log.Println("You don't have access")
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSession}
	} else {
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &status)
		if err != nil {
			log.Println("error json:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		} else {
			err = models.WareService.StatusWare(status.ID, status.Status)
			if err != nil {
				log.Println("ErrMysql:", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	}
	this.ServeJSON()
}
