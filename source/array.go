/**
 * @Author: Henry
 * @Description:
 * @File:  array.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 10:40 PM
 */

package main

import "fmt"

func main() {
	var str [2]string
	str[0] = "Hello"
	str[1] = "World"
	fmt.Println(str[0],str[1])
	fmt.Println(str)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}