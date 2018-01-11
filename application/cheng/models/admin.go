/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Inc.
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
 *     Initial: 2018/01/02        Shi Ruitao
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/shiruitao/go-one/application/cheng/utility"
)

type AdminServiceProvider struct {
}

var AdminService *AdminServiceProvider = &AdminServiceProvider{}

type Admin struct {
	Id       int16  `json:"id" orm:"pk"`
	Name     string `json:name`
	Password string `json:"password" validate:"required,alphanum,min=6,max=30"`
	State    int8   `json:"state"`
}

func init() {
	orm.RegisterModel(new(Admin))
}

func (create *AdminServiceProvider) Create(content Admin) error {
	content.State = 1
	o := orm.NewOrm()
	hash, err := utility.GenerateHash(content.Password)
	password := string(hash)
	sql := "INSERT INTO blog.admin(name, password) VALUES (?,?)"
	values := []interface{}{content.Name, password}
	raw := o.Raw(sql, values)
	_, err = raw.Exec()
	return err
}

func (login *AdminServiceProvider) Login(name, password string) (bool, error) {
	o := orm.NewOrm()
	var pwd string
	err := o.Raw("SELECT password FROM blog.admin WHERE name=? LIMIT 1 LOCK IN SHARE MODE", name).QueryRow(&pwd)
	if err != nil {
		return false, err
	} else if !utility.CompareHash([]byte(pwd), password) {
		return false, nil
	}
	return true, nil
}

func (change *AdminServiceProvider) ChangePass(name, newpassword string) error {
	o := orm.NewOrm()
	hash, err := utility.GenerateHash(newpassword)
	if err != nil {
		return err
	} else {
		password := string(hash)
		sql := "UPDATE blog.admin SET password = ? WHERE name = ? LIMIT 1"
		values := []interface{}{password, name}
		raw := o.Raw(sql, values)
		_, err := raw.Exec()

		return err
	}
}
