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
	"fmt"
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/cheng/log"
	"github.com/shiruitao/go-one/application/cheng/models"
)

// Operations about object
type Test struct {
	beego.Controller
}

func (t *Test) Insert() {
	var Content models.Message
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &Content)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"content": err}
	}
	id, err1 := models.MessageService.Insert(Content)
	fmt.Println(id, err1)
	if err1 != nil {
		log.Logger.Error("MessageService.Insert err %v:", err1)
		goto finish
	}
	t.Data["json"] = map[string]interface{}{"id": id}
finish:
	t.ServeJSON()
}

func (this *Test) ReadAll() {
	list, err := models.MessageService.ReadAll()
	if err != nil {
		log.Logger.Error("ERROR:", err)
	} else {
		this.Data["json"] = map[string]interface{}{"data": list}
	}
	this.ServeJSON()
}

func (t *Test) ReadLabel() {
	var label struct {
		Label string `json:"label"`
	}
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &label)
	fmt.Println("controllers:", t.Ctx.Input.RequestBody)
	if err != nil {
		log.Logger.Error("MessageService.Insert err %v:", err)
	}
	list := models.MessageService.ReadLabel(label.Label)
	t.Data["json"] = map[string]interface{}{"data": list}
	t.ServeJSON()
}

func (t *Test) ReadTitleContent() {
	var Mess struct {
		Title string `json:"title"`
	}
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &Mess)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
	}
	message, num := models.MessageService.ReadTitleContent(Mess.Title)
	t.Data["json"] = map[string]interface{}{"content:": message}
	if num == 0 {
		t.Data["json"] = map[string]string{"content:": "not find"}
	}
	fmt.Println("line:", num)
	t.ServeJSON()
}

var Id struct {
	Id []int `json:"id"`
}

func (this *Test) Delete() {
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Id)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
	}
	var (
		num  int64
		err1 error
	)
	for i := 0; i < len(Id.Id); i++ {
		num1, err1 := models.MessageService.Delete(Id.Id[i])
		if err1 != nil || num1 != 1 {
			log.Logger.Error("models.MessageService.Delete:", err1)
			break
		}
		num++
	}
	this.Data["json"] = map[string]interface{}{"line": num}
	if err1 != nil {
		log.Logger.Error("Delete(Id)", err1)
		goto finish
	}
finish:
	this.ServeJSON()
}

func (this *Test) DeleteTest() {
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Id)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
	}
	var num int64
	var err1 error
	for i := 0; i < len(Id.Id); i++ {
		num, err1 = models.MessageService.DeleteTest(Id.Id[i])
		if err1 != nil {
			break
		}
		num += 1
	}
	this.Data["json"] = map[string]interface{}{"执行数": num}
	this.ServeJSON()
}
