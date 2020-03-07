/**
 * @Author: Henry
 * @Description:
 * @File:  switch-evaluation-order.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 8:56 PM
 */

package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday{
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}
}
