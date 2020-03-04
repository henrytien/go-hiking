/**
 * @Author: HenryTien
 * @Description:
 * @File:  type-inference.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 10:40 AM
 */

package main

import	"fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T, value is of %v\n",v,v)
	f := 42.00
	fmt.Printf("f is of type %T, value is of %f\n",f,f)
}