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

type OrderServiceProvider struct{}

var OrderService OrderServiceProvider

type Order struct {
	ID      uint64    `orm:"column(id);pk;auto"`
	UserID  uint32    `orm:"column(userid)"`
	WareID  uint64    `orm:"column(wareid)" json:"ware_id"`
	Number  int8      `orm:"column(number)" json:"number"`
	Name    string    `orm:"column(name)" json:"name"`
	Address string    `orm:"column(address)" json:"address"`
	Phone   string    `orm:"column(phone)" json:"phone"`
	Finish  bool      `orm:"column(finish)"`
	Status  int8      `orm:"column(status)" json:"status"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)" json:"created"`
}

func init() {
	orm.RegisterModel(new(Order))
}

func (this *OrderServiceProvider) AddOrder(info *Order) (int64, error) {
	o := orm.NewOrm()

	order := Order{
		UserID:  info.UserID,
		WareID:  info.WareID,
		Number:  info.Number,
		Name:    info.Name,
		Address: info.Address,
		Phone:   info.Phone,
		Status:  1,
	}

	return o.Insert(&order)
}

func (this *OrderServiceProvider) FinishOrder(id uint64) (int64, error) {
	o := orm.NewOrm()

	order := Order{
		ID:     id,
		Status: 2,
	}

	return o.Update(order, "status")
}

func (this *OrderServiceProvider) DeleteOrder(id uint64) (int64, error) {
	o := orm.NewOrm()

	order := Order{
		ID:     id,
		Status: 3,
	}

	return o.Update(order, "status")
}
