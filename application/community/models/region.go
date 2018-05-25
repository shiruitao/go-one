package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type RegionServiceProvider struct{}

var RegionService *RegionServiceProvider

type Region struct {
	ID         uint32    `orm:"column(id);pk;auto"`
	Name       string    `orm:"column(name)" json:"name"`
	Address    string    `orm:"column(address)" json:"address"`
	Acreage    float32   `orm:"column(acreage)" json:"acreage"`
	Developers string    `orm:"column(developers)" json:"developers"`
	Created    time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Region))
}

func (*RegionServiceProvider) Create(info *Region) (int64, error) {
	var (
		r Region
	)

	r.Name = info.Name
	r.Acreage = info.Acreage
	r.Developers = info.Developers
	r.Address = info.Address
	o := orm.NewOrm()

	return o.Insert(&r)
}

func (*RegionServiceProvider) Get(name string) (*[]Region, int64, error) {
	var (
		r []Region
	)

	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM region WHERE name = ?", name).QueryRows(&r)
	return &r, num, err
}

func (*RegionServiceProvider) GetAll() (*[]Region, int64, error) {
	var (
		r []Region
	)

	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM region").QueryRows(&r)
	return &r, num, err
}

func (*RegionServiceProvider) Update(info *Region) (int64, error) {
	var (
		r Region
	)
	r.ID = info.ID
	r.Name = info.Name
	r.Address = info.Address
	r.Acreage = info.Acreage
	r.Developers = info.Developers

	o := orm.NewOrm()
	return o.Update(&r, "name", "address", "acreage", "developers")
}

func (*RegionServiceProvider) Delete(id uint32) (int64, error) {
	r := Region{ID: id}

	o := orm.NewOrm()
	return o.Delete(&r)
}
