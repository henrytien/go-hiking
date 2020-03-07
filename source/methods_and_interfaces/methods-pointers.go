/**
 * @Author: Henry
 * @Description:
 * @File:  methods-pointers.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 5:16 PM
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

func (v *Vertex) Scale(f float64) {
	 v.X = v.X * f
	 v.Y = v.Y * f
}

func main()  {
	v := Vertex{3, 4}
	v.Scale(104)
	fmt.Println(v.Abs())
}