/**
 * @Author: Henry
 * @Description:
 * @File:  defer-multi.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 9:54 PM
 */

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 1; i <= 5; i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")
}