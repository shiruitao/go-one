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
	//"log"
	"time"

	"github.com/astaxie/beego/orm"
	//"github.com/shiruitao/go-one/application/shop/utility"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type User struct {
	ID      uint32    `orm:"column(id);pk;auto"`
	Name    string    `orm:"column(name);null;utf8_bin" json:"name"`
	UnionID string    `orm:"column(unionid);unique" json:"union_id"`
	IsAdmin bool      `orm:"column(isadmin)" json:"is_admin"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func (this *UserServiceProvider) init() {
	o := orm.NewOrm()
	o.Using("store")
}

func (this *UserServiceProvider) CreateUser(u *User) (uint32, bool, error) {
	var (
		user User
		id   int64
		err  error
	)
	user.Name = u.Name
	user.UnionID = u.UnionID


	o := orm.NewOrm()
	if _, id, err = o.ReadOrCreate(&user, u.UnionID); err != nil {
		return uint32(id), false, err
	}
	user = User{ID: uint32(id)}
	err = o.Read(&user)
	return uint32(id), user.IsAdmin, err
}
