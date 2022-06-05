/**
 * @Author: Henry
 * @Description:
 * @File:  range.go
 * @Version: 1.0.0
 * @Date: 2020/3/6 3:35 PM
 */

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64}
func main() {
	for i , v := range pow{
		fmt.Printf("2**%d = %d\n",i ,v)
	}
}