/**
 * @Author: Henry
 * @Description:
 * @File:  type-assertions.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 2:01 PM
 */

package main

import (
	"fmt"
)

func main() {
	var i interface{} =  "mj"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s,ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)  // panic
	fmt.Println(f)

}