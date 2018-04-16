package main

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/shiruitao/go-one/application/shi/initorm"
	_ "github.com/shiruitao/go-one/application/shi/routers"
	"github.com/shiruitao/go-one/application/shi/models"
)

//func init() {
//	orm.RegisterModel(new(User1))
//}
//type User struct {
//	Id         int `orm:pk auto`
//	Int        int
//	Int8       int8
//	Int16      int16
//	Int32      int32
//	Int64      int64
//	Uint       uint
//	Uint16     uint16
//	Uint32     uint32
//	Uint64     uint64
//	Float32    float32
//	Float64    float64
//	Bool       bool
//	String     string
//	String128  string    `orm:"size(128)"`
//	StringText string    `orm:"type(text)"`
//	TimeDate   time.Time `orm:"type(date)"`
//	DateTime   time.Time
//	Byte       byte
//	Rune       rune
//}

func Table() {
	//orm.RunCommand()
	force := false
	verbose := true
	_ = orm.RunSyncdb("luoo", force, verbose)
}
func main() {
	orm.RegisterModel(new(models.User))
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	initorm.InitMysql()
	//beego.Run()
	Table()
}
