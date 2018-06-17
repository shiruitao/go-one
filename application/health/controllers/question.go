package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/health/models"
	"encoding/json"
)

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) Add() {
	var q models.Question

	json.Unmarshal(this.Ctx.Input.RequestBody, &q)

	models.QuestionService.Add(&q)
}
