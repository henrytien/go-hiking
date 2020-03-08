/**
 * @Author: Henry
 * @Description:
 * @File:  exercise-reader.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 4:05 PM
 */

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader)Read(s []byte)(int, error)  {
	s = s[:cap(s)]
	for i := range s{
		s[i] = 'A'
	}
	return cap(s),nil
}
func main()  {
	reader.Validate(MyReader{})
}