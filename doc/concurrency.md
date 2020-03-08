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

## Select

The `select` statement lets a goroutine wait on multiple communication operations.

A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
   x, y := 0, 1
   for {
      select {
      case c <- x:
         x, y = y, x+y
      case <-quit:
         fmt.Println("quit")
         return
      }
   }
}

func main() {
   c := make(chan int)
   quit := make(chan int)

   go func() {
      for i := 0; i < 10; i++ {
         fmt.Println(<-c)
      }
      quit <- 0
   }()
   fibonacci(c, quit)
}
```

Output :

```
0
1
1
2
3
5
8
13
21
34
quit
```

## Default Selection

The `default` case in a `select` is run if no other case is ready.

Use a `default` case to try a send or receive without blocking:

```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

```go
package main

import (
   "fmt"
   "time"
)

func main(){
   tick := time.Tick(1* time.Second)
   boom := time.Tick(5 *time.Second)
   for{
      select{
      case <- tick:
         fmt.Println("tick.")
      case <- boom:
         fmt.Println("boom.")
      default:
         fmt.Println("    mj, I love you")
         time.Sleep(3 * time.Second)
      }
   }
}
```

Output :

```
    mj, I love you
tick.
    mj, I love you
boom.
tick.
    mj, I love you
tick.
    mj, I love you
boom.
```

## Exercise: Equivalent Binary Trees

There can be many different binary trees with the same sequence of values stored in it. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

![img](https://tour.golang.org/content/img/tree.png)

A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

This example uses the `tree` package, which defines the type:

```
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```

**1.** Implement the `Walk` function.

**2.** Test the `Walk` function.

The function `tree.New(k)` constructs a randomly-structured (but always sorted) binary tree holding the values `k`, `2k`, `3k`, ..., `10k`.

Create a new channel `ch` and kick off the walker:

```
go Walk(tree.New(1), ch)
```

Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

**3.** Implement the `Same` function using `Walk` to determine whether `t1` and `t2` store the same values.

**4.** Test the `Same` function.

`Same(tree.New(1), tree.New(1))` should return true, and `Same(tree.New(1), tree.New(2))` should return false.

The documentation for `Tree` can be found [here](https://godoc.org/golang.org/x/tour/tree#Tree).

```go
package main

import (
   "fmt"
   "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
   WalkRecursive(t, ch)
   close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
   if t != nil {
      WalkRecursive(t.Left, ch)
      ch <- t.Value
      WalkRecursive(t.Right, ch)
   }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
   ch1, ch2 := make(chan int), make(chan int)
   go Walk(t1, ch1)
   go Walk(t2, ch2)
   for {
      n1, ok1 := <-ch1
      n2, ok2 := <-ch2
      if ok1 != ok2 || n1 != n2 {
         return false
      }

      if !ok1 {
         break
      }
   }
   return true
}

func main() {
   ch := make(chan int)
   go Walk(tree.New(1), ch)
   fmt.Println(Same(tree.New(1), tree.New(1)))
   fmt.Println(Same(tree.New(2), tree.New(1)))
   fmt.Println(Same(tree.New(1), tree.New(2)))
}
```

Output :

```
true
false
false
```

