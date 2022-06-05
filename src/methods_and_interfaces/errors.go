/**
 * @Author: Henry
 * @Description:
 * @File:  errors.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 3:19 PM
 */

package main

import (
	"fmt"
	"time"
)

type MyError struct{
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",e.When,e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work!",
	}
}

func main() {
	if err := run(); err != nil{
		fmt.Println(err)
	}
}