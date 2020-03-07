/**
 * @Author: Henry
 * @Description:
 * @File:  slices-literals.go
 * @Version: 1.0.0
 * @Date: 2020/3/5 10:35 AM
 */

package main

import "fmt"

func main() {
	q := []int{3, 7, 9, 11, 13}
	fmt.Println(q)

	r := []bool{false, true, false, true}
	fmt.Println(r)

	s := []struct{
		x int
		b bool
	}{
		{1, false},
		{2, false},
		{520, true},
	}
	fmt.Println(s)
}