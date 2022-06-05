/**
 * @Author: Henry
 * @Description:
 * @File:  exercise-loops-and-functions.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 8:09 PM
 */

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i:= 1; i <= 10; i++{
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	p :=2.0
	fmt.Println(Sqrt(p))
	fmt.Println(math.Sqrt(p))
}
