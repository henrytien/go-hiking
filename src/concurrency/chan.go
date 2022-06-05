/**
 * @Author: Henry
 * @Description:
 * @File:  chan.go
 * @Version: 1.0.0
 * @Date: 1/5/21 3:39 PM
 */

package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA()chan int {
	ch := make(chan int, 10)
	// 启动一个goroutine 用于生成随机数，函数返回一个通道用于获取随机数
	go func() {
		for {
			ch <-rand.Int()
		}
	}()
	return ch
}
func main() {
	ch := GenerateIntA()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}