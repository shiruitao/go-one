package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type HistoryServiceProvider struct{}

var HistoryService *HistoryServiceProvider

type HistoryData struct {
	ID           uint32    `orm:"column(id);pk;auto"`
	DeviceID     string    `orm:"column(DviceID)" json:"deviceID"`
	DataTimeHour time.Time `orm:"column(DataTimeHour)" json:"dataTimeHour"`
	CurrentA     float32   `orm:"column(CurrentA)" json:"currentA"`
	CurrentB     float32   `orm:"column(CurrentB)" json:"currentB"`
	CurrentC     float32   `orm:"column(CurrentC)" json:"currentC"`
	VoltA        float32   `orm:"column(VoltA)" json:"voltA"`
	VoltB        float32   `orm:"column(VoltB)" json:"voltB"`
	VoltC        float32   `orm:"column(VoltC)" json:"voltC"`
	IsAvailable  bool      `orm:"column(isAvailable)" json:"isAvailable"`
	IsCalculate  bool      `orm:"column(isCalculate)" json:"isCalculate"`
	EnergyA      float32   `orm:"column(EnergyA)" json:"energyA"`
	EnergyB      float32   `orm:"column(EnergyB)" json:"energyB"`
	EnergyC      float32   `orm:"column(EnergyC)" json:"energyC"`
}

func init() {
	orm.RegisterModel(new(HistoryData))
}

func (*HistoryServiceProvider) Get() ([]HistoryData, error) {
	var historyData []HistoryData

	o := orm.NewOrm()
	qs := o.QueryTable("historydata")
	_, err := qs.All(&historyData)

	return historyData, err
}

//func (*HistoryServiceProvider) GetTime() ([]HistoryData, error) {
//	var historyData []HistoryData
//
//	o := orm.NewOrm()
//	qs := o.QueryTable("historydata")
//	_, err := qs.Filter("DataTimeHour", t).All(&historyData)
//
//	return historyData, err
//}
