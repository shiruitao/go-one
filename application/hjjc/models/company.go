package models

import (
	"github.com/astaxie/beego/orm"
)

type CompanyServiceProvider struct{}

var CompanyService *CompanyServiceProvider

type Company struct {
	ID   uint32  `orm:"column(id);pk;auto"`
	Name string  `orm:"column(name)" json:"name"`
	Area string  `orm:"column(divisions)" json:"divisions"`
	JD   float64 `orm:"column(jd)" json:"jd"`
	WD   float64 `orm:"column(wd)" json:"wd"`
}

func init() {
	orm.RegisterModel(new(Company))
}

func (*CompanyServiceProvider) Add(info *Company) error {
	var c Company

	c.Name = info.Name
	c.Area = info.Area
	c.JD = info.JD
	c.WD = info.WD

	o := orm.NewOrm()
	_, err := o.Insert(&c)
	return err
}

func (*CompanyServiceProvider) Get() ([]Company, error) {
	var company []Company

	o := orm.NewOrm()
	qs := o.QueryTable("company")
	_, err := qs.All(&company)

	return company, err
}

func (*CompanyServiceProvider) Update(info *Company) error {
	var c Company

	c.ID = info.ID
	c.Name = info.Name
	c.Area = info.Area
	c.JD = info.JD
	c.WD = info.WD

	o := orm.NewOrm()
	_, err := o.Update(&c)

	return err
}

func (*CompanyServiceProvider) Delete(id uint32) error {
	c := Company{ID: id}

	o := orm.NewOrm()
	_, err := o.Delete(&c)
	return err
}
