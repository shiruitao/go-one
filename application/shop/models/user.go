/*
 * MIT License
 *
 * Copyright (c) 2018 Shi Ruitao.
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
 *     Initial: 2018/03/10        Shi Ruitao
 */

package models

import (
	"log"

	"github.com/astaxie/beego/orm"

	"github.com/shiruitao/go-one/application/shop/utility"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type User struct {
	ID       uint32 `orm:"column(id)"pk auto`
	Name     string `orm:"column(name)"`
	Phone    string `orm:"column(phone)"`
	PassWord string `orm:"column(pass_word)"`
	State    int8   `orm:"column(state)"`
}

type (
	CreateUser struct {
		Name        string `json:"name" validate:"required`
		PassWord    string `json:"pass_word" validate:"required,min=6,max=30"`
		ConfirmPass string `json:"confirm_pass"`
	}
)

func (this *UserServiceProvider) init() {
	o := orm.NewOrm()
	o.Using("default")
	//orm.RegisterModel(new(User))
}

func (this *UserServiceProvider) CreateUser(u *CreateUser) (int64, error) {
	var id int64

	hash, err := utility.GenerateHash(u.PassWord)
	if err != nil {
		log.Println(err)
		return id, err
	}
	password := string(hash)

	o := orm.NewOrm()
	var user User
	user.Name = u.Name
	user.PassWord = password
	id, err = o.Insert(&user)

	return id, err
}
