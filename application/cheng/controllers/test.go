package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/cheng/log"
	"github.com/shiruitao/go-one/application/cheng/models"
	//"github.com/astaxie/beego/orm"
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

var Mess struct {
	Title string `json:"title"`
	Content string `json:"content"`
}
func (t *Test) ReadTitleContent() {
	fmt.Println("123456")
	//err := json.Unmarshal(t.Ctx.Input.RequestBody, &Mess)
	//if err != nil {
	//	log.Logger.Error("json.Unmarshal:", err)
	//}
	models.MessageService.ReadTitleContent("JavaScript")
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
			fmt.Println("执行数", num1)
			log.Logger.Error("models.MessageService.Delete:", err1)
			break
		}
		num++
	}
	fmt.Println("执行", num, "行")
	this.Data["json"] = map[string]interface{}{"执行数": num}
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

func (t *Test) Update() {
	models.MessageService.Update(1001)
}
