/**
 * @Author: Henry
 * @Description:
 * @File:  range-continued.go
 * @Version: 1.0.0
 * @Date: 2020/3/6 3:48 PM
 */

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println(pow)
	// index
	for i := range pow{
		pow[i] = 1 << uint(i) // == 2**i
	}
	// value
	for _, value := range pow{
		fmt.Printf("%d\n", value)
	}
}