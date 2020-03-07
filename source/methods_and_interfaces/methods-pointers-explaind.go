/**
 * @Author: Henry
 * @Description:
 * @File:  methods-pointers-explained.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 5:33 PM
 */

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X , Y float64
}

func Abs(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3,4}
	Scale(&v, 4)
	fmt.Println(Abs(v))
}