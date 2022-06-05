/**
 * @Author: Henry
 * @Description:
 * @File:  channel_chain.go
 * @Version: 1.0.0
 * @Date: 1/5/21 9:12 PM
 */

package main

import "fmt"

// chain 函数的输入参数和输出参数类型相同，都是chan  int 类型
// chain 函数的功能是将chan内的数据统一加1
func chain(in chan int) chan int{
	out := make (chan int)
	go func() {
		for v :=range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int)
	// 初始化输入参数
	go func() {
		for i:= 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	// 连续调用3次chan，相当于把in中的每个元素都加3
	out := chain(chain(chain(in)))
	for v:= range out {
		fmt.Println(v)
	}
}