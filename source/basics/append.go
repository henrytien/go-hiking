/**
 * @Author: Henry
 * @Description:
 * @File:  append.go
 * @Version: 1.0.0
 * @Date: 2020/3/5 10:18 PM
 */

package main

import "fmt"

func main() {
	var s[]int
	printSlices(s)

	// append works on nil slices
	s = append(s, 0)
	printSlices(s)

	// the slices grows as need
	s = append(s, 1)
	printSlices(s)

	// we can add more elements at a time
	s = append(s, 520,1,1314)
	printSlices(s)
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n",len(s), cap(s), s)
}