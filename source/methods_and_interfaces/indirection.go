/**
 * @Author: Henry
 * @Description:
 * @File:  indirection.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 5:42 PM
 */

package main

import "fmt"

type Vertex struct {
	X ,Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScalFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3,4}
	v.Scale(4)
	ScalFunc(&v,4)

	fmt.Println(v)

	p := &Vertex{3,4}
	p.Scale(4)
	ScalFunc(p,4)
	fmt.Println(p)
	fmt.Println(*p)

}