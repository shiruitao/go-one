package models

import (
	"time"
	"github.com/astaxie/beego/orm"

)

type EconomicServiceProvider struct{}

var EconomicService *EconomicServiceProvider

type Economic struct {
	ID      uint32    `orm:"column(id);pk;auto"`
	Year    int       `orm:"column(year)" json:"year"`
	count   int       `orm:"column(count)" json:"count"`
	Area    string    `orm:"column(area)" json:"area"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func (*EconomicServiceProvider) Create(info *Economic) (int64, error) {
	var (
		e Economic
	)
	e.Year = info.Year
	e.Area = info.Area
	e.count = info.count

	o := orm.NewOrm()
	return o.Insert(&e)
}

func (*EconomicServiceProvider) Get(area string) (*Economic, int64, error) {
	var (
		economic Economic
	)

	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM economic WHERE area = ? ORDER BY id DESC", area).QueryRows(&economic)
	return &economic, num, err
}
