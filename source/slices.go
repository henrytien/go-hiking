/**
 * @Author: Henry
 * @Description:
 * @File:  slices.go
 * @Version: 1.0.0
 * @Date: 2020/3/5 9:53 AM
 */

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}