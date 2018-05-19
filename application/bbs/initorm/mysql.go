package initorm

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func InitMysql() {
	host := beego.AppConfig.String("mysql::host")
	port, _ := beego.AppConfig.Int("mysql::port")
	user := beego.AppConfig.String("mysql::user")
	pass := beego.AppConfig.String("mysql::pass")
	db := beego.AppConfig.String("mysql::db")
	maxIdle, _ := beego.AppConfig.Int("mysql::idle")
	maxConn, _ := beego.AppConfig.Int("mysql::conn")

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, host, port, db), maxIdle, maxConn)
	orm.RegisterDataBase("bbs", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, host, port, "bbs"), maxIdle, maxConn)
}