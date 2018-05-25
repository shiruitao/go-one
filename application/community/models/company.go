package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type CompanyServiceProvider struct{}

var CompanyService *CompanyServiceProvider

type Company struct {
	ID          uint32    `orm:"column(id);pk;auto"`
	Name        string    `orm:"column(name)" json:"name"`
	Type        string    `orm:"column(type)" json:"type"`
	Capital     int       `orm:"column(capital)" json:"capital"`
	ManageScope string    `orm:"column(managescope)" json:"scope"`
	Address     string    `orm:"column(address)" json:"address"`
	Area        string    `orm:"column(area)" json:"area"`
	Created     time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Company))
}

func (*CompanyServiceProvider) Create(info *Company) error {
	var (
		company Company
	)

	company = Company{
		Name: info.Name,
		Type: info.Type,
		Capital: info.Capital,
		ManageScope: info.ManageScope,
		Address: info.Address,
		Area: info.Area,
	}

	o := orm.NewOrm()

	_, err := o.Insert(&company)
	if err != nil {
		return err
	}
	return nil
}

func (*CompanyServiceProvider) GetAll() (*[]Company, int64, error) {
	var (
		company []Company
	)

	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM company").QueryRows(&company)
	return &company, num, err
}

func (*CompanyServiceProvider) GetByArea(area string) (*[]Company, int64, error) {
	var (
		company []Company
	)

	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM company WHERE area = ? ORDER BY id DESC", area).QueryRows(&company)
	return &company, num, err
}

func (*CompanyServiceProvider) Delete(id uint32) error {
	company := Company{ID: id}

	o := orm.NewOrm()
	_, err := o.Delete(&company)
	return err
}
