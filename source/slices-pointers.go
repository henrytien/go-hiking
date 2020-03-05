/**
 * @Author: Henry
 * @Description:
 * @File:  slices-pointers.go
 * @Version: 1.0.0
 * @Date: 2020/3/5 10:29 AM
 */

package main

import "fmt"

func main(){
	names := [4]string {
		"mj",
		"henry",
		"lbq",
		"jy",
	}

	fmt.Println(names)
	a := names[0:1]
	fmt.Println(a)
	b := names[1:4]
	fmt.Println(b)
	b[0] = "henry20"
	fmt.Println(b)
	fmt.Println(names)
}