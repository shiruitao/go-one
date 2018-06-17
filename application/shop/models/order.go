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
	"github.com/astaxie/beego/orm"
)

type OrderServiceProvider struct{}

var OrderService OrderServiceProvider

type Order struct {
	ID      uint64    `orm:"column(id);pk;auto"`
	Title   string	  `orm:"column(title)" json:"title"`
	Number  int8      `orm:"column(num)" json:"num"`
	Name    string    `orm:"column(name)" json:"name"`
	Price    float32    `orm:"column(price)" json:"price"`
	image   string  `orm:"column(image)" json:"image"`
	Address string    `orm:"column(address)" json:"address"`
	Phone   string    `orm:"column(phone)" json:"phone"`
}

func init() {
	orm.RegisterModel(new(Order))
}

func (this *OrderServiceProvider) AddOrder(info *Order) (int64, error) {
	o := orm.NewOrm()

	order := Order{
		Number:  info.Number,
		Name:    info.Name,
		Address: info.Address,
		Phone:   info.Phone,
	}

	return o.Insert(&order)
}

func (this *OrderServiceProvider) FinishOrder(id uint64) (int64, error) {
	o := orm.NewOrm()

	order := Order{
		ID:     id,
	}

	return o.Update(order, "status")
}

func (this *OrderServiceProvider) DeleteOrder(id uint64) (int64, error) {
	o := orm.NewOrm()

	order := Order{
		ID:     id,
	}

	return o.Delete(&order)
}

func (this *OrderServiceProvider) GetOrder() (*[]Order, error) {
	var (
		order []Order
	)

	o := orm.NewOrm()
	_, err := o.QueryTable("order").All(&order)
	return &order, err
}
