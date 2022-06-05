/**
 * @Author: Henry
 * @Description:
 * @File:  stringer.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 2:25 PM
 */

package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func (p Person)String() string{
	return fmt.Sprintf("%v (%v years)",p.Name, p.Age)
}

func main() {
	m := Person{"mj", 18}
	h := Person{"henry",25}
	fmt.Println(m,h)
}