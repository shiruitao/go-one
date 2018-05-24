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
 *     Initial: 2018/05/24        Shi Ruitao
 */

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type IndustryServiceProvider struct{}

var IndustryService *IndustryServiceProvider

type Industry struct {
	ID          uint32    `orm:"column(id);pk;auto"`
	Manufacture int       `orm:"column(manufacture)" json:"manufacture"`
	Build       int       `orm:"column(build)" json:"build"`
	Retail      int       `orm:"column(retail)" json:"retail"`
	Catering    int       `orm:"column(catering)" json:"catering"`
	Financial   int       `orm:"column(financial)" json:"financial"`
	Internet    int       `orm:"column(internet)" json:"internet"`
	Area        string    `orm:"column(area)" json:"area"`
	Created     time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func (*IndustryServiceProvider) Get(area string) (*[]Industry, int64, error) {
	var (
		industry []Industry
	)

	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM industry WHERE area = ? ORDER BY id DESC", area).QueryRows(&industry)
	return &industry, num, err
}
