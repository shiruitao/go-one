package models

import (
	"github.com/astaxie/beego/orm"
)

type CompanyServiceProvider struct {}

var CompanyService *CompanyServiceProvider

type Company struct {
	ID uint32 `orm:"column(id);pk;auto"`
	Name string `orm:"column(name)" json:"name"`
	Area string `orm:"column(area)" json:"area"`
	JD float64 `orm:"column(jd)"`
	WD float64 `orm:"column(wd)"`
}

func init() {
	orm.RegisterModel(new(Company))
}

func (*CompanyServiceProvider) Get(info *Company) (*[]Company, error) {
	var company []Company

	o := orm.NewOrm()
	qs := o.QueryTable("company")
	_, err := qs.Filter("area", info.Area).Filter("name", info.Name).All(&company)

	return &company, err
}
