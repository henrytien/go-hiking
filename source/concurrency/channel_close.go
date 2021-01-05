/**
 * @Author: Henry
 * @Description:
 * @File:  channel_close.go
 * @Version: 1.0.0
 * @Date: 1/5/21 8:09 PM
 */

package main

import (
	"fmt"
)

func GenerateIntA(done chan struct{}) chan int{
	ch:= make(chan int)
	go func() {
		Lable:
			for {
				// 通过select监听一个信号 chan来确定是否停止生成
				select {
				//case ch <- rand.Int():
				case <-done:
					break Lable
				}
			}
			close(ch)
	}()
	return ch
}

func main() {
	done :=  make(chan struct{

	})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 不再需要生成器，通过close chan发送一个通知给生成器
	close(done)
	for v := range ch {
		fmt.Println(v)
	}

}