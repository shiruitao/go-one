package models

import (
	"github.com/astaxie/beego/orm"
)

type ImageServiceProvider struct{}

var ImageService *ImageServiceProvider

type Image struct {
	Image string `form:"image" orm:"pk"`
}

func init() {
	orm.RegisterModel(new(Image))
}

func (insert *ImageServiceProvider) Insert(image string) {
}
