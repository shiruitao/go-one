package models

import (
	"github.com/astaxie/beego/orm"
)
type DevicesServiceProvider struct{}

var DevicesService *DevicesServiceProvider

type Devices struct {
	ID            uint32  `orm:"column(id);pk;auto"`
	DeviceID      uint32  `orm:"column(DviceID)" json:"deviceID"`
	DeviceName    string  `orm:"column(DeviceName)" json:"deviceName"`
	SetupDate     string  `orm:"column(SetupDate)" json:"setupDate"`
	DivisionID    uint32  `orm:"column(DivisionID)" json:"divisionID"`
	DeviceAddr    string  `orm:"column(DeviceAddr)" json:"deviceAddr"`
	CompanyName   string  `orm:"column(CompanyName)" json:"companyName"`
	DeviceType    string  `orm:"column(DeviceType)" json:"deviceType"`
	LastDataTime  string  `orm:"column(LastDataTime)" json:"lastDataTime"`
	CompanyID     uint32  `orm:"column(CompanyID)" json:"companyID"`
	PhaseACurrent float32 `orm:"column(PhaseACurrent)" json:"phaseACurrent"`
	PhaseBCurrent float32 `orm:"column(PhaseBCurrent)" json:"phaseBCurrent"`
	PhaseCCurrent float32 `orm:"column(PhaseCCurrent)" json:"phaseCCurrent"`
	CTRatio       float32 `orm:"column(CTRatio)" json:"CTRatio"`
}

func init() {
	orm.RegisterModel(new(Devices))
}

func (*DevicesServiceProvider) Add(info *Devices) error {
	var d Devices

	d.DeviceID = info.DeviceID
	d.DeviceName = info.CompanyName
	d.SetupDate = info.SetupDate
	d.DivisionID = info.DivisionID
	d.DeviceAddr = info.DeviceAddr
	d.CompanyName = info.CompanyName
	d.DeviceType = info.DeviceType
	d.LastDataTime = info.LastDataTime
	d.CompanyID = info.CompanyID
	d.PhaseACurrent = info.PhaseACurrent
	d.PhaseBCurrent = info.PhaseBCurrent
	d.PhaseCCurrent = info.PhaseCCurrent
	d.CTRatio = info.CTRatio

	o := orm.NewOrm()
	_, err := o.Insert(&d)
	return err
}

func (*DevicesServiceProvider) Get() ([]Devices, error) {
	var device []Devices

	o := orm.NewOrm()
	qs := o.QueryTable("devices")
	_, err := qs.All(&device)

	return device, err
}

func (*DevicesServiceProvider) Update(info *Devices) error {
	var d Devices

	d.ID = info.ID
	d.CompanyID = info.CompanyID
	d.DeviceName = info.CompanyName
	d.SetupDate = info.SetupDate
	d.DivisionID = info.DivisionID
	d.DeviceAddr = info.DeviceAddr
	d.CompanyName = info.CompanyName
	d.DeviceType = info.DeviceType
	d.LastDataTime = info.LastDataTime
	d.CompanyID = info.CompanyID
	d.PhaseACurrent = info.PhaseACurrent
	d.PhaseBCurrent = info.PhaseBCurrent
	d.PhaseCCurrent = info.PhaseCCurrent
	d.CTRatio = info.CTRatio

	o := orm.NewOrm()
	_, err := o.Update(&d)

	return err
}

func (*DevicesServiceProvider) Delete(id uint32) error {
	d := Devices{ID: id}

	o := orm.NewOrm()
	_, err := o.Delete(&d)
	return err
}

