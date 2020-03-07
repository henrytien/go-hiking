/**
 * @Author: Henry
 * @Description:
 * @File:  interfaces.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 10:45 PM
 */

package main

import (
	"fmt"
	"math"
)

type Abser interface{
	Abs() float64
}

func main(){
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}
	fmt.Println(f)
	a = f  // a MyFloat implements Abser

	a = &v // a *Vertex implements Abser

	// In the flowing line, v is a Vertex(not *Vertex)
	// and does not implement Abser
	a = &v
	fmt.Println(a.Abs())

}



type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

