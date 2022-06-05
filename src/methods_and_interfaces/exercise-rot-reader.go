/**
 * @Author: Henry
 * @Description:
 * @File:  exercise-rot-reader.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 4:46 PM
 */

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct{
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	n, e := r.r.Read(b)
	if e == nil{
		for i,v := range b{
			switch  {
			case v >= 'A' && v <= 'Z':
				b[i] = (v - 'A' + 13) % 26 + 'A'
			case v >= 'a' && v <='z':
				b[i] = (v - 'a' + 13) % 26 + 'a'
			}
		}
	}
	return
}

func main(){
	s := strings.NewReader("HELLO, mj")
	r := rot13Reader{s}
	io.Copy(os.Stdout,&r)
}