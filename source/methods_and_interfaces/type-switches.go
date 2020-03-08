/**
 * @Author: Henry
 * @Description:
 * @File:  type-switches.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 2:11 PM
 */

package main

import "fmt"

func do(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("Third %v is %v\n", v, v*3)
	case string:
		fmt.Printf("%q is %v bytes long\n",v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

func main(){
	do("mj")
	do(520)
	do(true)
}