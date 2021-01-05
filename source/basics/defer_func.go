/**
 * @Author: Henry
 * @Description:
 * @File:  defer_func.go
 * @Version: 1.0.0
 * @Date: 1/5/21 3:20 PM
 */

package main

func f1() (r int) {

	defer func() {
		r ++
	}()
	return
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r  = r + 5
	}(r)
	return
}


func main() {
	println("f1=",f1())
	println("f2=",f2())
	println("f3=",f3())
}
//f1= 1
//f2= 5
//f3= 0
