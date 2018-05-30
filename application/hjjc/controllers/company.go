package controllers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/hjjc/common"
	"github.com/shiruitao/go-one/application/hjjc/log"
	"github.com/shiruitao/go-one/application/hjjc/models"
)

type CompanyController struct {
	beego.Controller
}

func (this *CompanyController) Get() {
	type position struct{
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	}

	pos := []position{}

	info, err := models.CompanyService.Get()
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		for i := 0; i <100; i++ {
			pos[i].Latitude = info[i].WD
			pos[i].Longitude = info[i].JD
			log.Logger.Info("", info[i])
		}
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": info, "position": pos}
	}

	this.ServeJSON()
}
