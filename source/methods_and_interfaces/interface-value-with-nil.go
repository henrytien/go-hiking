/**
 * @Author: Henry
 * @Description:
 * @File:  interface-value-with-nil.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 1:12 PM
 */

package main

import (
	"fmt"
)

type I interface {
	M()
}
type T struct {
	S string
}
func (t *T) M()  {
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	i = t
	describe(t)
	i.M()

	t = &T{"mj"}
	t.M()
}

func describe(i I)  {
	fmt.Printf("(%v, %T)\n",i, i)
}
