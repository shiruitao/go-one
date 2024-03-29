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

package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/cheng/common"
	"github.com/shiruitao/go-one/application/cheng/log"
	"github.com/shiruitao/go-one/application/cheng/models"
)

type Test struct {
	beego.Controller
}

func (t *Test) Insert() {
	var Content models.Message
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &Content)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"content": err}
	} else {
		err = models.MessageService.Insert(Content)
		if err != nil {
			log.Logger.Error("MessageService.Insert err %v:", err)
		} else {
			t.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	t.ServeJSON()
}

func (this *Test) ReadAll() {
	list, num, err := models.MessageService.ReadAll()
	if err != nil {
		log.Logger.Error("ERROR:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{"line": num, common.RespKeyData: list}
	}
	this.ServeJSON()
}

func (t *Test) ReadLabel() {
	var label struct {
		Label1 string `json:"label1"`
		Label2 string `json:"label2"`
	}
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &label)
	if err != nil {
		log.Logger.Error("MessageService.Insert err %v:", err)
	} else {
		list, num, err := models.MessageService.ReadLabel(label.Label1, label.Label2)
		if err != nil {
			log.Logger.Error("Readlabel error:", err)
		} else {
			t.Data["json"] = map[string]interface{}{"line": num, common.RespKeyData: list}
		}
	}
	t.ServeJSON()
}

func (this *Test) ReadTitleContent() {
	var Mess struct {
		Title string `json:"title"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Mess)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
	}
	message, num, err := models.MessageService.ReadTitleContent(Mess.Title)
	this.Data["json"] = map[string]interface{}{"content:": message}
	if err != nil {
		log.Logger.Error("ReadTitleContent error:", err)
		this.Data["json"] = map[string]string{"content:": "not find"}
	} else {
		this.Data["json"] = map[string]interface{}{"line:": num, common.RespKeyData: message}
	}
	this.ServeJSON()
}

func (this *Test) ReadTime() {
	var Date struct {
		Date string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Date)
	if err != nil {
		log.Logger.Error("time error", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: "time error"}
	} else {
		messages, num, err := models.MessageService.ReadTime(Date.Date)
		if err != nil {
			log.Logger.Error("time read error", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		} else {
			this.Data["json"] = map[string]interface{}{"line": num, common.RespKeyData: messages}
		}
	}
	this.ServeJSON()
}

func (this *Test) Delete() {
	var Id struct {
		Id []int `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Id.Id)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
	} else {
		var (
			err1 error
		)
		err1 = models.MessageService.Delete(&Id.Id)
		if err1 != nil {
			log.Logger.Error("models.MessageService.Delete:", err1)
			this.Data["json"] = map[string]interface{}{"status": "ErrMysql"}
		}

		this.Data["json"] = map[string]interface{}{"status": "success"}
		if err1 != nil {
			log.Logger.Error("Delete(Id)", err1)
		}
	}
	this.ServeJSON()
}

func (this *Test) DeleteTest() {
	var Id struct {
		Id []int `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Id)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
	}
	_, err = models.MessageService.DeleteTest(&Id.Id)
	if err != nil {
		log.Logger.Error("mysqlErr:", err)
		this.Data["json"] = map[string]interface{}{"status": "ErrMysql"}
	}
	this.Data["json"] = map[string]interface{}{"status": "success"}
	this.ServeJSON()
}

func (this *Test) Modify() {
	var (
		mess models.Message
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &mess)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
	}
	err = models.MessageService.Modify(&mess)
	if err != nil {
		log.Logger.Error("models.MessageService.Modify:", err)
		this.Data["json"] = map[string]interface{}{"status": "ErrMysql"}
	}
	this.Data["json"] = map[string]interface{}{"status": "success"}
	this.ServeJSON()
}
