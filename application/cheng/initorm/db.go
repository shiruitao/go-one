package initorm

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
	// 参数1   driverName
	// 参数2   数据库类型
	// 这个用来设置 driverName 对应的数据库类型
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, host, port, db), maxIdle, maxConn)
	orm.RegisterDataBase("blog", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, host, port, "blog"), maxIdle, maxConn)
}
