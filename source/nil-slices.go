/**
 * @Author: Henry
 * @Description:
 * @File:  nil-slices.go
 * @Version: 1.0.0
 * @Date: 2020/3/5 11:31 AM
 */

package main

import "fmt"

func main() {
	var s[] int
	if len(s) == 0{
		fmt.Println("0")
	}

	fmt.Println(s,len(s),cap(s))

	if s == nil{
		fmt.Println("nil!")
	}
	// panic: runtime error: index out of range [0] with length 0
	s[0] = 3
	fmt.Println(s)
}