/**
 * @Author: Henry
 * @Description:
 * @File:  interface-values.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 12:56 PM
 */

package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T)M()  {
	fmt.Println(t.S)
}

type F float64

func (f F)M()  {
	fmt.Println(f*10)
}

func main() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I)  {
	fmt.Printf("(%v, %T)\n", i, i)
}