/**
 * @Author: Henry
 * @Description:
 * @File:  struct-fileds.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 10:21 PM
 */

package main

import "fmt"

type Vertex struct{
	X int
	Y int
}

func main() {
	v := Vertex{1,4}
	fmt.Println(v)
	v.X = 5
	v.Y = 0
	fmt.Println(v)
}