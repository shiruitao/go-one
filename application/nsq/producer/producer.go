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
 *     Initial: 2018/05/11        Shi Ruitao
 */

package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

// 主函数
func main() {
	strIP1 := "127.0.0.1:4150"
	strIP2 := "127.0.0.1:4152"
	InitProducer(strIP1)

	running := true

	//读取控制台输入
	// reader := bufio.NewReader(os.Stdin)
	index := 1
	for running {
		// data, _, _ := reader.ReadLine()

		command := fmt.Sprintf("I'm number %d !", index)
		fmt.Println(command)
		index++
		if command == "stop" {
			running = false
		}

		var err error
		for err = Publish("test", command); err != nil; err = Publish("test", command) {
			//切换IP重连
			fmt.Println("换线了!!!")
			strIP1, strIP2 = strIP2, strIP1
			InitProducer(strIP1)
		}
		time.Sleep(time.Second)
	}
	//关闭
	producer.Stop()
}

// 初始化生产者
func InitProducer(str string) {
	var err error
	fmt.Println("address: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
}

//发布消息
func Publish(topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic, []byte(message)) // 发布消息
		return err
	}
	return fmt.Errorf("producer is nil", err)
}
