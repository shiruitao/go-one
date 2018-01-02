package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	"github.com/shiruitao/go-one/application/cheng/models"
	"fmt"
)

// Operations about object
type Test struct {
	beego.Controller
}

func (t *Test) Insert() {
	var Content models.Message
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &Content)
	id, err := models.MessageService.Insert(Content)
	fmt.Println(id, err)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"content": err}
		goto finish
	}
	t.Data["json"] = map[string]interface{}{"id": id}
finish:
	t.ServeJSON()
}

func (t *Test) Read() {
	err := models.MessageService.Read(1004)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"content": err}
		goto finish
	}
	finish:
	t.ServeJSON()
}

func (this *Test) Delete() {
	var (
		num int64
		id struct {
			Id int `json:"id"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		fmt.Println("controller中:", this.Ctx.Input.RequestBody)
		fmt.Println("err:", err)
		this.Data["json"] = map[string]interface{}{"content": err}
		goto finish
	}
	
	num, err = models.MessageService.Delete(id.Id)
	fmt.Println("controller中2:", id.Id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"content": err, "num": num}
		goto finish
	} else {
		this.Data["json"] = map[string]interface{}{"执行行数": num}
	}
	finish:
	this.ServeJSON()
}

func (t *Test) Update() {
	models.MessageService.Update(1001)
}