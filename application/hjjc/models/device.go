package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type DeviceServiceProvider struct{}

var DeviceService *DeviceServiceProvider

type (
	Device struct {
		ID       uint32  `orm:"column(id);pk;auto"`
		Name     string  `orm:"column(name)" json:"name"`
		Device   string  `orm:"column(device)" json:"device"`
		DataTime string  `orm:"column(datatime)" json:"data_time"`
		Aa       float32 `orm:"column(Aa)" json:"aa"`
		Ab       float32 `orm:"column(Ab)" json:"ab"`
		Ac       float32 `orm:"column(Ac)" json:"ac"`
		Va       float32 `orm:"column(Va)" json:"va"`
		Vb       float32 `orm:"column(Vb)" json:"vb"`
		Vc       float32 `orm:"column(Vc)" json:"vc"`
		Wa       float32 `orm:"column(Wa)" json:"wa"`
		Wb       float32 `orm:"column(Wb)" json:"wb"`
		Wc       float32 `orm:"column(Wc)" json:"wc"`
		Area     string  `orm:"column(area)" json:"area"`
		Created  time.Time `orm:"column(created)" json:"created"`
	}
	ReqInfo struct {
		Area     string `json:"area"`
		Name  string `json:"name"`
		Device   string `json:"device"`
		DataTime string `json:"data_time"`
	}
)

func init() {
	orm.RegisterModel(new(Device))
}

func (*DeviceServiceProvider) Get(info *ReqInfo) (*[]Device, error) {
	var (
		device []Device
	)
	o := orm.NewOrm()
	qs := o.QueryTable("device")
	_, err := qs.Filter("area", info.Area).Filter("name", info.Name).Filter("device", info.Device).Filter("datatime", info.DataTime).All(&device)

	return &device, err
}
