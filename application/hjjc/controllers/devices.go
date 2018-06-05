package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	
	"github.com/shiruitao/go-one/application/hjjc/models"
	"github.com/shiruitao/go-one/application/hjjc/log"
	"github.com/shiruitao/go-one/application/hjjc/common"
)

type DevicesController struct {
	beego.Controller
}

func (this *DevicesController) Get() {
	info, err := models.DevicesService.Get()
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": info}
	}

	this.ServeJSON()
}

func (this *DevicesController) Add() {
	var devices models.Devices

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &devices)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.DevicesService.Add(&devices)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *DevicesController) Delete() {
	var id struct{
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err = models.DevicesService.Delete(id.ID)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
		this.ServeJSON()
	}

}

func (this *DevicesController) Update() {
	var devices models.Devices

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &devices)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.DevicesService.Update(&devices)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}