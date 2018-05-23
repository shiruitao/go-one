/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/05/22        Shi Ruitao
 */

package initorm

import (
	"github.com/astaxie/beego"
	"fmt"
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

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, host, port, db), maxIdle, maxConn)
	orm.RegisterDataBase("community", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, host, port, "community"), maxIdle, maxConn)
}