/**
 * @Author: Henry
 * @Description:
 * @File:  simple_generator.go
 * @Version: 1.0.0
 * @Date: 1/8/21 8:13 PM
 */

package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA() chan int {
	ch := make(chan int,10)
	// 启动一个goroutine 用于生成随机数，函数返回一个通道用于获取随机数
	go func() {
		for   {
			ch<-rand.Int()
		}
	}()
	return ch
}

func GenerateIntB() chan int {
	ch := make(chan int,10)
	// 启动一个goroutine 用于生成随机数，函数返回一个通道用于获取随机数
	go func() {
		for   {
			ch<-rand.Int()
		}
	}()
	return ch
}

func GenerateInt() chan int {
	ch := make(chan int,20)
	// 启动一个goroutine 用于生成随机数，函数返回一个通道用于获取随机数
	go func() {
		for   {
			// 使用select的扇入技术(Fan in) 增加随机生成的随机源
			select {
			case ch<- <-GenerateIntA():
			case ch<- <-GenerateIntB():
			}
		}
	}()
	return ch
}

func main() {
	ch := GenerateInt()
	for i:= 0; i < 100; i++  {
		fmt.Println(<-ch)
	}
	fmt.Println(<-ch)
}