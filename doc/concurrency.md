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

