/**
 * @Author: Henry
 * @Description:
 * @File:  for.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 11:04 AM
 */

package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("Total is %v", sum)
}