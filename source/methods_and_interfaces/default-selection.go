/**
 * @Author: Henry
 * @Description:
 * @File:  default-selection.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 10:38 PM
 */

package main

import (
	"fmt"
	"time"
)

func main(){
	tick := time.Tick(1* time.Second)
	boom := time.Tick(5 *time.Second)
	for{
		select{
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("boom.")
		default:
			fmt.Println("    mj, I love you")
			time.Sleep(3 * time.Second)
		}
	}
}