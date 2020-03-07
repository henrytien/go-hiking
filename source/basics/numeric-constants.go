/**
 * @Author: Henry
 * @Description:
 * @File:  numeric-constants.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 10:53 AM
 */

package main

import "fmt"

const(
	// Create a huger number by shifting a 1 bit left 100 places
	// In other words, the binary number is 1 followed by 100 zeros.
	Big = 1 << 100
	// Shift it right again 99 places, so we end with 1 << 1 or 2.
	Small = Big >> 99
)

func needInt(x int) int { return 10*x + 1}
func needFloat(x float64) float64 { return x*0.1}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	//fmt.Println(needInt(Big))  //overflows int
	fmt.Println(needFloat(Big))
}