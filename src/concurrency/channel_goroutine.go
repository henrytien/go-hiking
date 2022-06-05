/**
 * @Author: Henry
 * @Description:
 * @File:  channel_goroutine.go
 * @Version: 1.0.0
 * @Date: 1/5/21 8:46 PM
 */

package main

import (
	"fmt"
	"math/rand"
)

// done 接收通知退出信号
func GenenerateIntA(done chan  struct{}) chan int  {
	ch := make(chan int,5)

	go func() {
		Lable:
			for{
				select {
				case ch<-rand.Int():
					case <-done:
						break Lable
				}
			}
			close(ch)
	}()
	return ch
}

// done 接收通知退出信号
func GenerateIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
		Lable:
			for{
				select {
					case ch <- rand.Int():
						case <-done:
							break Lable
				}
			}
			close(ch)
	}()
	return ch
}

// 通过select 执行扇入（Fan in)操作
func GenerateInt(done chan struct{}) chan int{
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
		Lable:
			for{
				select {
				case ch <- <- GenenerateIntA(send):
				case ch <- <-GenerateIntB(send):
				case <-done:
					send <- struct{}{}
					send <- struct{}{}
					break Lable
				}
			}
			close(ch)
	}()
	return ch
}

func main() {
	// 创建一个作为接收退出的信号的chan
	done := make (chan struct{})
	// 	启动生成器
	ch := GenerateInt(done)

	// 获取生成器资源
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// 通知生产者停止生产
	done <- struct{}{}
	fmt.Println("stop gernerate")
}