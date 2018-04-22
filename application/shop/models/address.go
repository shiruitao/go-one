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
 *     Initial: 2018/04/22        Shi Ruitao
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AddressServiceProvider struct{}

var AddressService AddressServiceProvider

type Address struct {
	ID        uint32    `orm:"column(id);pk;auto" json:"id"`
	UserID    uint32    `orm:"column(userid)" json:"user_id"`
	Name      string    `orm:"column(name)" json:"name"`
	Address   string    `orm:"column(address)" json:"address"`
	Phone     string    `orm:"column(phone)" json:"phone"`
	IsDefault bool      `orm:"column(isdefault)" json:"is_default"`
	Created   time.Time `orm:"column(created);auto_now_add;type(datetime)" json:"created"`
}

func init() {
	orm.RegisterModel(new(Address))
}

func (this *AddressServiceProvider) AddAddress(info *Address) error {
	o := orm.NewOrm()
	addr := Address{
		UserID:    info.UserID,
		Name:      info.Name,
		Address:   info.Address,
		Phone:     info.Phone,
		IsDefault: info.IsDefault,
	}
	_, err := o.Insert(&addr)
	return err
}
