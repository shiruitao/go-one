package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/cheng/models"
	"fmt"
)

// Operations about object
type Test struct {
	beego.Controller
}

func (o *Test) Helloworld() {
	o.Data["json"] = map[string]string{"content": models.Helloworld("shiruitao")}
	o.ServeJSON()
}
func (add *Test) Add() {
	add.Data["json"] = models.Add(3, 5)
}

func (t *Test) Insert() {
	id, err := models.MessageService.Insert()
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
	err := models.MessageService.Read()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"content": err}
		goto finish
	}
finish:
	t.ServeJSON()
}
