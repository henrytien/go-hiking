/**
 * @Author: Henry
 * @Description:
 * @File:  if-and-else.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 7:43 PM
 */

package main

import (
	"fmt"
	"math"
)

func pow(x , n , lim float64) float64  {
	if v:= math.Pow(x,n); v < lim{
		return v
	}else{
		fmt.Printf("%g >= %g\n",v,lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(2,6,40),
		pow(2,5,520),
		)
}