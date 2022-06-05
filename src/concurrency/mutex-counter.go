/**
 * @Author: Henry
 * @Description:
 * @File:  mutex-counter.go
 * @Version: 1.0.0
 * @Date: 2020/3/9 11:18 AM
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct{
	v map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key
func (c *SafeCounter)Inc(key string){
	c.mux.Lock()
	// Lock so only one goroutine can access the map c.v
	c.v[key]+= 2
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter)Get(key string) int{
	c.mux.Lock()
	// Lock so only one goroutine can access the map c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v:make(map[string]int)}
	for i := 0; i < 5; i++{
		go c.Inc("mj")
		go c.Inc("henry")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Get("mj"),c.Get("henry"))

}