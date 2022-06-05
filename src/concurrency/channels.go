/**
 * @Author: Henry
 * @Description:
 * @File:  channels.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 8:28 PM
 */

package main

import "fmt"

func sum(s []int,c chan int){
	sum := 0
	for _,v := range s{
		sum += v 
	}
	c <- sum
}

func main()  {
	s := []int{3,4,5,2,0}

	c := make(chan int)

	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	 x := <- c
	 y := <- c
	 fmt.Println(x,y)
	 fmt.Println(x + y)
}