package controllers

import (
	"github.com/shiruitao/go-one/application/cheng/models"
	"github.com/astaxie/beego"
)

// Operations about object
type Test struct {
	beego.Controller
}
func (o *Test) Helloworld() {
	o.Data["json"] = map[string]string{"content": models.Helloworld("shiruitao")}
	o.ServeJSON()
}
