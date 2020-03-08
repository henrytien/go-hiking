/**
 * @Author: Henry
 * @Description:
 * @File:  goroutines.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 7:45 PM
 */

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++{
		time.Sleep(2 *time.Second)
		fmt.Println(s)
	}
}

func main() {
	go say("jy")
	say("I love u")
}