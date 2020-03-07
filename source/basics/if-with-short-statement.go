/**
 * @Author: Henry
 * @Description:
 * @File:  if-with-short-statement.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 11:48 AM
 */

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x,n); v < lim{
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(2,8,520),
		pow(3,6,520),
		)
}