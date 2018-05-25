package models

import (
	"time"
)

type RegionServiceProvider struct {}

var RegionService *RegionServiceProvider

type Region struct {
	ID        uint32    `orm:"column(id);pk;auto"`
	Name      string    `orm:"column(name)" json:"name"`
	Address   string    `orm:"column(address)" json:"address"`
	Area      string    `orm:"column(area)" json:"area"`
	Created   time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}


