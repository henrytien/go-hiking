/**
 * @Author: henry
 * @Description:
 * @File:  variable-with-initializers.go
 * @Version: 1.0.0
 * @Date: 2020/3/3 9:54 PM
 */

package main

import "fmt"

var i, j int = 3, 4

func main() {
	var python, c, java  = true, false ,"yes"
	fmt.Println(i,j,python,c,java)
}

