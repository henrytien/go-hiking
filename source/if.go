/**
 * @Author: Henry
 * @Description:
 * @File:  if.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 11:43 AM
 */

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return  fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(520),sqrt(-520))
}