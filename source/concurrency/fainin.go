/**
 * @Author: Henry
 * @Description:
 * @File:  fainin.go
 * @Version: 1.0.0
 * @Date: 1/8/21 7:50 PM
 */

package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// GenerateIntA 是一个随机数生成器
func GenerateIntA(done chan struct{}) chan int {
	ch:= make(chan int)
	go func() {
		Lable:
			for{
				select {
					case ch <- rand.Int():
					// 增加一路监听，就是退出的通知信号done的监听
					case <- done:
						break Lable
				}
			}
			close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 发送通知，告诉生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 此时生产者已经退出
	println("NumGoroutine=", runtime.NumGoroutine())
	
}