package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/cheng/common"
	"github.com/shiruitao/go-one/application/cheng/log"
	"github.com/shiruitao/go-one/application/cheng/models"
)

type ImageController struct {
	beego.Controller
}

func (this *ImageController) Image() {
	image := models.Image{}
	if err := this.ParseForm(&image); err != nil {
		log.Logger.Error("image error:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyData: image}
	}
	this.ServeJSON()
}
