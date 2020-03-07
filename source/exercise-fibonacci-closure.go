/**
 * @Author: Henry
 * @Description:
 * @File:  exercise-fibonacci-closure.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 2:42 PM
 */

package main

import "fmt"

// fibonacci is a function that returns
// a function that return an int.
func fibonacci() func() int {
	a, b := 0,1
	return func() int {
		a,b = b, a + b
		return a
	}
}

func main() {
	f := fibonacci()
	for i:= 0; i < 10; i++{
		fmt.Println(f())
	}
}