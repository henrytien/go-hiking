/**
 * @Author: Henry
 * @Description:
 * @File:  maps.go
 * @Version: 1.0.0
 * @Date: 2020/3/6 4:44 PM
 */

package main

import "fmt"

type Vetex struct{
	Lat, Long float64
}

var m map[string]Vetex

func main() {
	m = make(map[string]Vetex)
	m["mj"] = Vetex{
		Lat: 2,
		Long: 0,
	}
	fmt.Println(m)
	fmt.Println(m["mj"])
}