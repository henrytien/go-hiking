/**
 * @Author: henry
 * @Description:
 * @File:  named-results.go
 * @Version: 1.0.0
 * @Date: 2020/3/3 9:44 PM
 */

package main

import "fmt"
func split(sum int) (x, y int) {
	x = sum * 5 / 10
	y = sum - x
	return
}

func main() {
	fmt.Println(split(30))
}