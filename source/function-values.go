/**
 * @Author: Henry
 * @Description:
 * @File:  function-values.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 11:13 AM
 */

package main

import (
	"fmt"
	"math"
)

func compute(fn func(x, y float64) float64) float64 {
	return fn(4, 3)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))

	fmt.Println(compute(math.Pow))
}