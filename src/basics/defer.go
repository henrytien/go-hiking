/**
 * @Author: Henry
 * @Description:
 * @File:  defer.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 9:29 PM
 */

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}