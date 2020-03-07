/**
 * @Author: Henry
 * @Description:
 * @File:  methods-continued.go
 * @Version: 1.0.0
 * @Date: 2020/3/7 4:36 PM
 */

package main

import (
	"fmt"
	"math"
)

type Myfloat float64

func (f Myfloat)Abs()float64  {
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}

func main()  {
	f := Myfloat(-math.SqrtE)
	fmt.Println(f.Abs())
}