package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type IndustryServiceProvider struct{}

var IndustryService *IndustryServiceProvider

type Industry struct {
	ID          uint32    `orm:"column(id);pk;auto"`
	Manufacture int       `orm:"column(manufacture)" json:"manufacture"`
	Build       int       `orm:"column(build)" json:"build"`
	Retail      int       `orm:"column(retail)" json:"retail"`
	Catering    int       `orm:"column(catering)" json:"catering"`
	Financial   int       `orm:"column(financial)" json:"financial"`
	Internet    int       `orm:"column(internet)" json:"internet"`
	Area        string    `orm:"column(area)" json:"area"`
	Created     time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func (*IndustryServiceProvider) Create(info *Industry) (int64, error) {
	var (
		i Industry
	)
	i.Manufacture = info.Manufacture
	i.Build = info.Build
	i.Retail = info.Retail
	i.Catering = info.Catering
	i.Financial = info.Financial
	i.Internet = info.Internet
	i.Area = info.Area

	o := orm.NewOrm()

	return o.Insert(&i)
}

func (*IndustryServiceProvider) Get(area string) (*[]Industry, int64, error) {
	var (
		industry []Industry
	)

	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM industry WHERE area = ? ORDER BY id DESC", area).QueryRows(&industry)
	return &industry, num, err
}
