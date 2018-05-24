package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/community/common"
	"github.com/shiruitao/go-one/application/community/log"
	"github.com/shiruitao/go-one/application/community/models"
)

type CompanyController struct {
	beego.Controller
}

func (this *CompanyController) CreateCompany() {
	var company models.Company

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &company)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err = models.CompanyService.Create(&company)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *CompanyController) GetCompanyAll() {
	company, num, err := models.CompanyService.GetAll()
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": company, "number": num}
	}
	this.ServeJSON()
}

func (this *CompanyController) GetCompanyByArea() {
	var area struct {
		Area string `json:"area"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &area)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		company, num, err := models.CompanyService.GetByArea(area.Area)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": company, "number": num}
		}
	}
	this.ServeJSON()
}

func (this *CompanyController) DeleteCompany() {
	var id struct {
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.CompanyService.Delete(id.ID)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}
