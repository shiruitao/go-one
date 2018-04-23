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
 *     Initial: 2018/04/23        Shi Ruitao
 */

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type CartServiceProvider struct{}

var CartService CartServiceProvider

type Cart struct {
	ID      uint64    `orm:"column(id);pk;auto" json:"id"`
	UserID  uint32    `orm:"column(uerid)" json:"user_id"`
	WareID  string    `orm:"column(wareid)" json:"ware_id"`
	Number  int8      `orm:"column(number)" json:"number"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)" json:"created"`
}

func init() {
	orm.RegisterModel(new(Cart))
}

func (this *CartServiceProvider) AddCart(info *Cart) (int64, error) {
	o := orm.NewOrm()
	cart := Cart{
		UserID: info.UserID,
		WareID: info.WareID,
		Number: info.Number,
	}

	return o.Insert(&cart)
}

func (this *CartServiceProvider) ModifyNum(id uint64, num int8) (int64, error) {
	o := orm.NewOrm()

	cart := Cart{
		ID:     id,
		Number: num,
	}

	return o.Update(&cart, "number")
}

func (this *CartServiceProvider) DeleteWare(id uint64) (int64, error) {
	o := orm.NewOrm()

	cart := Cart{ID: id}

	return o.Delete(&cart)
}
