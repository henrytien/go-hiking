/**
 * @Author: Henry
 * @Description:
 * @File:  images.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 7:10 PM
 */

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
}