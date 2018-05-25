package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type HouseServiceProvider struct{}

var HouseService *HouseServiceProvider

type House struct {
	ID         uint32    `orm:"column(id);pk;auto"`
	Name       string    `orm:"column(name)" json:"name"`
	Layer      int       `orm:"column(layer)" json:"layer"`
	High       float32   `orm:"column(high)" json:"high"`
	Address    string    `orm:"column(address)" json:"address"`
	Acreage    float32   `orm:"column(acreage)" json:"acreage"`
	Date       string    `orm:"column(date)" json:"builddate"`
	Region     string    `orm:"column(region)" json:"region"`
	Created    time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(House))
}

func (*HouseServiceProvider) Create(info *House) (int64, error) {
	var (
		h House
	)

	h.Name = info.Name
	h.Layer = info.Layer
	h.High = info.High
	h.Date = info.Date
	h.Region = info.Region
	h.Acreage = info.Acreage
	h.Address = info.Address
	o := orm.NewOrm()

	return o.Insert(&h)
}

func (*HouseServiceProvider) GetAll() (*[]House, error) {
	var (
		h []House
	)

	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM house").QueryRows(&h)
	return &h, err
}

func (*HouseServiceProvider) GetByRegion(name string) (*[]House, error) {
	var (
		h []House
	)

	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM house WHERE region = ?", name).QueryRows(&h)
	return &h, err
}

func (*HouseServiceProvider) Update(info *House) (int64, error) {
	var (
		h House
	)
	h.ID = info.ID
	h.Name = info.Name
	h.Address = info.Address
	h.Acreage = info.Acreage
	h.Region = info.Region
	h.High = info.High
	h.Date = info.Date
	h.Layer = info.Layer

	o := orm.NewOrm()
	return o.Update(&h, "name", "address", "acreage", "region", "high", "date", "layer")
}

func (*HouseServiceProvider) Delete(id uint32) (int64, error) {
	h := House{ID: id}
	o := orm.NewOrm()

	return o.Delete(&h)
}
