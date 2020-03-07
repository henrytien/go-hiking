/**
 * @Author: henry
 * @Description:
 * @File:  multiple-results.go
 * @Version: 1.0.0
 * @Date: 2020/3/3 9:34 PM
 */

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main()  {
	//fmt.Println(swap("hello", "world"))
	a,b := swap("Hello", "World")
	fmt.Println(a, b)
}