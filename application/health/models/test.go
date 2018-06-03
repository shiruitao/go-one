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
 *     Initial: 2018/06/03        Shi Ruitao
 */

package models

import (
	"github.com/astaxie/beego/orm"
)

type TestServiceProvider struct{}

var TestService *TestServiceProvider

type Test struct {
	ID    uint32 `orm:"column(id);pk;auto"`
	Title uint32 `orm:"column(title)" json:"title"`
	A     int    `orm:"column(A)" json:"a"`
	B     int    `orm:"column(B)" json:"b"`
	C     int    `orm:"column(C)" json:"c"`
	D     int    `orm:"column(D)" json:"d"`
}

func init() {
	orm.RegisterModel(new(Test))
}

func (*TestServiceProvider) Get(id uint32) (*Test, error) {
	var test Test

	o := orm.NewOrm()

	err := o.QueryTable("test").Filter("title", id).One(&test)

	return &test, err
}
