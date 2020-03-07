/**
 * @Author: Henry
 * @Description:
 * @File:  indirection-values.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 5:55 PM
 */

package main

import "fmt"

type Vertex struct {
	X , Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v Vertex, f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

func main() {
	v := Vertex{3,4}
	v.Scale(4)
	fmt.Println(v)
	fmt.Println(ScaleFunc(v,4))

	p := &Vertex{3,4}
	p.Scale(4)
	fmt.Println(p)
	fmt.Println(ScaleFunc(*p,4))

}