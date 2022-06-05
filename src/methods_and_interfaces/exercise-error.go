/**
 * @Author: Henry
 * @Description:
 * @File:  exercise-error.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 3:32 PM
 */

package main

import (
	"fmt"
	"math"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not sqrt negative number %f", e)
}
const e = 1e-8
func Sqrt(x float64) (float64,error) {
	if x < 0{
		return 0,ErrNegativeSqrt(x)
	}

	z := x
	for{
		newZ := z - ((z*z) - x) / (z*2)
		if math.Abs(newZ- z) < e{
			return newZ,nil
		}
		z = newZ
	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}