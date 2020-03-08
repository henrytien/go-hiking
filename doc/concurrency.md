## Goroutines

A *goroutine* is a lightweight thread managed by the Go runtime.

```
go f(x, y, z)
```

starts a new goroutine running

```
f(x, y, z)
```

The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The [`sync`](https://golang.org/pkg/sync/) package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)

```go
package main

import (
   "fmt"
   "time"
)

func say(s string) {
   for i := 0; i < 5; i++{
      time.Sleep(2 *time.Second)
      fmt.Println(s)
   }
}

func main() {
   go say("mj")
   say("I love u")
}
```

Output :

```
I love u
mj
I love u
mj
I love u
mj
I love u
mj
mj
I love u
```

## Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

```
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

```go
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
```

Output :

```
7 7
14
```

## Buffered Channels

Channels can be *buffered*. Provide the buffer length as the second argument to `make` to initialize a buffered channel:

```
ch := make(chan int, 100)
```

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 5
	ch <- 20
	//ch <- 5  // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	//fmt.Println(<- ch)  //fatal error: all goroutines are asleep - deadlock!
}
```

Output :

```
5
20
```

## Range and Close

A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```
v, ok := <-ch
```

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

**Note:** Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

**Another note:** Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.

```go
package main

import "fmt"

func fibonacci(n int,c chan int) int {
	x, y := 0, 1
	for i := 1; i < n; i++{
		c <- i
		x , y = y, x + y
	}
	//fmt.Println("s")
	close(c)
	return y
}

func main() {
	c := make(chan int,10)
	go fibonacci(cap(c), c)
	//fmt.Println(fibonacci(cap(c),c))
	for i := range c{
		fmt.Println(i)
	}
}
```

Output :

```
1
2
3
4
5
6
7
8
9
```

