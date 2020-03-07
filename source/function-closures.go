/**
 * @Author: Henry
 * @Description:
 * @File:  function-closures.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 11:28 AM
 */

package main

import "fmt"

func adder() func(int)int {
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}

func main() {
	pos , neg := adder(), adder()
	for i:= 0; i < 10; i++{
		fmt.Println(pos(i),
			neg(i))
	}
}