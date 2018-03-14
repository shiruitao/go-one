package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserServerProvider struct {
}

var UserServer *UserServerProvider

type User struct {
	Id         int `orm:pk auto`
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Float32    float32
	Float64    float64
	Bool       bool
	String     string
	String128  string `orm:size 128`
	StringText string `orm:type(text)`
	Time       time.Time
	Byte       byte
	Rune       rune
}

func (create *UserServerProvider) init() {
	o := orm.NewOrm()
	o.Using("default")
	//orm.RegisterModel(new(User))
}

func (create *UserServerProvider) Create() {
}
