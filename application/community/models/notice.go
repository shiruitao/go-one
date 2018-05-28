package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type NoticeServiceProvider struct{}

var NoticeService *NoticeServiceProvider

type Notice struct {
	ID      uint32    `orm:"column(id);pk;auto"`
	Title   string    `orm:"column(title)" json:"title"`
	Content string    `orm:"column(content)" json:"content"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Notice))
}

func (*NoticeServiceProvider) Create(info *Notice) (int64, error) {
	var (
		n Notice
	)
	n.Title = info.Title
	n.Content = info.Content

	o := orm.NewOrm()
	return o.Insert(&n)
}

func (*NoticeServiceProvider) Get() (*[]Notice, error) {
	var n []Notice

	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM notice ORDER BY id DESC").QueryRows(&n)
	return &n, err
}
