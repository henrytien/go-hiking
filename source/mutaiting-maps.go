/**
 * @Author: Henry
 * @Description:
 * @File:  mutaiting-maps.go
 * @Version: 1.0.0
 * @Date: 2020/3/6 5:13 PM
 */

package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["mj"] = 520
	fmt.Println("The value: ", m["mj"])

	m["mj"] = 20
	fmt.Println("The value: ",m["mj"])

	delete(m,"mj")
	fmt.Println("The value: ",m["mj"])

	v, ok := m["mj"]
	fmt.Println("The value ", v, "Present?", ok)

}