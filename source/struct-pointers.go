/**
 * @Author: Henry
 * @Description:
 * @File:  struct-pointers.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 10:26 PM
 */

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1,10}
	fmt.Println(v)
	p := &v
	(*p).X = 10
	p.Y = 100
	fmt.Println(v)
}