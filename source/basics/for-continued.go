/**
 * @Author: Henry
 * @Description:
 * @File:  for-continued.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 11:28 AM
 */

package main

import "fmt"

func main() {
	sum := 1
	for ;sum < 1000;  {
		sum += sum
		//fmt.Println(sum)
	}
	fmt.Println(sum)
}