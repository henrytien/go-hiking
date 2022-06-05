/**
 * @Author: Henry
 * @Description:
 * @File:  empty-interface.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 1:52 PM
 */

package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 2
	describe(i)

	i = "mj"
	describe(i)
}

func describe(i interface{})  {
	fmt.Printf("(%v, %T)\n",i ,i)
}