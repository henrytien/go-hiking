/**
 * @Author: Henry
 * @Description:
 * @File:  reader.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 3:51 PM
 */

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello Reader!")
	b := make([]byte, 8)
	for{
		n, err := r.Read(b)
		fmt.Println(n)
		fmt.Printf("n = %v err = %v b = %v\n",r, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF{
			break
		}
	}
}