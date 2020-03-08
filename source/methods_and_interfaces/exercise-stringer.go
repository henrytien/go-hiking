/**
 * @Author: Henry
 * @Description:
 * @File:  exercise-stringer.go
 * @Version: 1.0.0
 * @Date: 2020/3/8 2:35 PM
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

//func (ip IPAddr) String() string {
//	return fmt.Sprintf("%d.%d.%d.%d",ip[0],ip[1],ip[2],ip[3])
//}

func (ip IPAddr) String() string {
	s := make([]string,len(ip))
	for i, val := range ip{
		s[i] = strconv.Itoa(int(val))
	}
	return strings.Join(s,".")

}

func main() {
	hosts := map[string]IPAddr{
		"loopback":{127,0,0,1},
		"googleDNS":{8,8,8,8},
	}

	for name,ip := range hosts{
		fmt.Printf("%v: %v\n",name,ip)
	}

}