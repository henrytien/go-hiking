/**
 * @Author: Henry
 * @Description:
 * @File:  buffered-channels.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 9:26 PM
 */

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 5
	ch <- 20
	//ch <- 5  // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	//fmt.Println(<- ch)  //fatal error: all goroutines are asleep - deadlock!
}