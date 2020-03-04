/**
 * @Author: Henry
 * @Description:
 * @File:  struct-literals.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 10:32 PM
 */

package main

import "fmt"

type Vertex struct{
	X, Y int
}

var(

	v1 = Vertex{X:1}		// Y:0 is implicit
	v2 = Vertex{}			// X:0 and Y:0
	p = &Vertex{1,3}	// has type *Vertex
)

func main() {
	v := Vertex{1,2}	// has type Vertex
	fmt.Println(v, v1, v2, p, *p)
}