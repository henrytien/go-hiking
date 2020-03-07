/**
 * @Author: Henry
 * @Description:
 * @File:  map-literals.go
 * @Version: 1.0.0
 * @Date: 2020/3/6 4:50 PM
 */

package main

import "fmt"

type Info struct{
	Name string
	Id int
}

var m = map[string]Info{
	"mj":Info{
		"mj", 101,
	},
	"henry":Info{"henry",100,
	},
}

func main()  {
	fmt.Println(m)
}