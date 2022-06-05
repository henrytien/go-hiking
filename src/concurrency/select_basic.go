/**
 * @Author: Henry
 * @Description:
 * @File:  select_basic.go
 * @Version: 1.0.0
 * @Date: 1/8/21 7:38 PM
 */

package main

func main() {
	ch :=  make(chan int,1)
	go func(chan int) {
		for  {
			select {
			// 0 或 1的写入是随机的
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)

	for i := 0; i < 10; i++ {
		println(<-ch)
	}
}