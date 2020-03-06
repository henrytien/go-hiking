/**
 * @Author: Henry
 * @Description:
 * @File:  map-literal-continued.go
 * @Version: 1.0.0
 * @Date: 2020/3/6 5:05 PM
 */

package main

import "fmt"

type Info struct{
	Id int
	Name string
}

var m = map[string]Info{
	"henry":{1 ,"henry"},
	"mj":{2,"mj"},
}

func main() {
	fmt.Println(m)
}