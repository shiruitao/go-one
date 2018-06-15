package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/bbs/common"
	"github.com/shiruitao/go-one/application/bbs/log"
	"github.com/shiruitao/go-one/application/bbs/models"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ArticleCreate() {
	var (
		article models.Article
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		article.AuthorID = this.GetSession(common.SessionUserID).(uint32)
		err = models.ArticleService.New(&article)
		if err != nil {
			log.Logger.Error("ErrDB:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *ArticleController) DeleteArt() {
	var id struct {
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleService.DeleteArt(id.ID)
		if err != nil {
			log.Logger.Error("ErrDB:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
}

func (this *ArticleController) DeleteUser() {
	var id struct {
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleService.DeleteUser(id.ID)
		if err != nil {
			log.Logger.Error("ErrDB:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

}

func (this *ArticleController) Get() {
	article, err := models.ArticleService.Get()
	if err != nil {
		log.Logger.Error("ErrDB:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": article}
	}
	this.ServeJSON()
}
