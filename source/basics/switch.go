/**
 * @Author: Henry
 * @Description:
 * @File:  switch.go
 * @Version: 1.0.0
 * @Date: 2020/3/4 8:43 PM
 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Printf("%s.\n",os)
	}
}
