/**
 * @Author: henry
 * @Description:
 * @File:  type-conversions.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 10:28 AM
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 6, 9
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var u uint = uint(f)
	fmt.Println(x, y, f, u)
}
