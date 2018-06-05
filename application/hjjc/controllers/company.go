package controllers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/hjjc/common"
	"github.com/shiruitao/go-one/application/hjjc/log"
	"github.com/shiruitao/go-one/application/hjjc/models"
	"encoding/json"
)

type CompanyController struct {
	beego.Controller
}

func (this *CompanyController) Get() {
	info, err := models.CompanyService.Get()
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": info}
	}

	this.ServeJSON()
}

func (this *CompanyController) Add() {
	var company models.Company

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &company)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.CompanyService.Add(&company)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *CompanyController) Delete() {
	var id struct{
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err = models.CompanyService.Delete(id.ID)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
		this.ServeJSON()
	}

}

func (this *CompanyController) Update() {
	var company models.Company

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &company)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.CompanyService.Update(&company)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}
