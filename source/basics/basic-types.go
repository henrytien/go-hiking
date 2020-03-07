/**
 * @Author: henry
 * @Description:
 * @File:  basic-types.go
 * @Version: 1.0.0
 * @Date: 2020/3/3 10:35 PM
 */

package main

import (
	"fmt"
	"math/cmplx"
)
var	(
	Tobe bool = false
	MaxInt uint64 = 1 << 64-1
	z  complex128 = cmplx.Sqrt(-5 + 12i)
)
func main() {
	fmt.Printf("Type: %T Value: %v\n",Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n",MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n",z, z)
}