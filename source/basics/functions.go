/**
 * @Author: henry
 * @Description:
 * @File:  functions.go
 * @Version: 1.0.0
 * @Date: 2020/3/3 9:22 PM
 */

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main(){
	fmt.Println(add(3,5))
}