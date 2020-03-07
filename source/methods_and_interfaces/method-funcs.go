/**
 * @Author: Henry
 * @Description:
 * @File:  method-funcs.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 4:28 PM
 */

package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

func Abs(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{5,20}
	fmt.Println(Abs(v))
}