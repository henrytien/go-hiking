/**
 * @Author: Henry
 * @Description:
 * @File:  slice-bounds.go
 * @Version: 1.0.0
 * @Date: 2020/3/5 11:02 AM
 */

package main

import "fmt"

func main() {
	s := []int{2, 3, 520 ,11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}