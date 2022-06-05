/**
 * @Author: Henry
 * @Description:
 * @File:  methods.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 3:07 PM
 */

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v:= Vertex{2,4}
	fmt.Println(v.Abs)
	fmt.Println(v.Abs())
}
