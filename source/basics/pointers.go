/**
 * @Author: Henry
 * @Description:
 * @File:  pointers.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 10:00 PM
 */

package main

import "fmt"

func main() {
	i, j := 4, 5

	p := &i			// point to i
	fmt.Println(*p) // read i through the pointer
	*p = 32 		// set i through the pointer
	fmt.Println("the new value of i ",i)  // see the new value of i

	p = &j   		// point to j
	*p = *p / 23 	// divide j through the pointer
	fmt.Println("the new value of j ", j)  // see the new value of j

}