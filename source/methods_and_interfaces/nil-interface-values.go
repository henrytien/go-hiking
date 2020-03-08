/**
 * @Author: Henry
 * @Description:
 * @File:  nil-interface-values.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 1:24 PM
 */

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M()  {
	//if t == nil{
	//	return
	//}
	fmt.Println(t.S)
}

func main() {
	var i I
	describe(i)
	var t *T
	describe(t)
	i = t
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)", i ,i)
}