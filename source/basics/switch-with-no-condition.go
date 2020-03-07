/**
 * @Author: Henry
 * @Description:
 * @File:  switch-with-no-condition.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 9:05 PM
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Hour(),t.Minute(),t.Second())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
