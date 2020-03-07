/**
 * @Author: Henry
 * @Description:
 * @File:  interfaces-are-satisfied-implicitly.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 11:02 PM
 */

package main

import "fmt"

type I interface{
	M()
}

type T struct{
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so
func (t T) M()  {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"mj"}
	i.M()
	i = T{"i love u"}
	i.M()
}


