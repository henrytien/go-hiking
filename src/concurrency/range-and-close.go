/**
 * @Author: Henry
 * @Description:
 * @File:  range-and-close.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 9:36 PM
 */

package main

import "fmt"

func fibonacci(n int,c chan int) int {
	x, y := 0, 1
	for i := 1; i < n; i++{
		c <- i
		x , y = y, x + y
	}
	//fmt.Println("s")
	close(c)
	return y
}

func main() {
	c := make(chan int,10)
	go fibonacci(cap(c), c)
	//fmt.Println(fibonacci(cap(c),c))
	for i := range c{
		fmt.Println(i)
	}
}