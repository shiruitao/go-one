package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/hjjc/models"
)

type AbnormalController struct {
	beego.Controller
}

func (this *AbnormalController) Get() {
	var (
		res, res1 []float32
		num = 0
	)
	info, _ := models.DevicesService.Get()
	historyData, _ := models.HistoryService.Get()
	for i, value := range info {
		res[i] = value.PhaseACurrent
	}
	for i, value := range historyData {
		res1[i] = value.CurrentB
	}
	for i := 0; i < len(res); i++ {
		if res1[i] < res[i] {
			num ++
		}
	}
	this.Data["json"] = map[string]interface{}{"abnormal": num, "companyNum": len(res)}
	this.ServeJSON()
}
