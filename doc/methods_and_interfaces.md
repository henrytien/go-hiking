## Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special *receiver* argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` named `v`.

```go
package main

import (
   "fmt"
   "math"
)

type Vertex struct {
   X, Y float64
}

func (v Vertex) Abs() float64 {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
   v:= Vertex{2,4}
   fmt.Println(v.Abs)
   fmt.Println(v.Abs())
}
```

Output :

```
0x109cfb0
4.47213595499958
```

## Methods are functions

Remember: a method is just a function with a receiver argument.

Here's `Abs` written as a regular function with no change in functionality.

```go
package main

import (
   "fmt"
   "math"
)

type Vertex struct{
   X, Y float64
}

func Abs(v Vertex) float64{
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
   v := Vertex{5,20}
   fmt.Println(Abs(v))
}
```

Output :

```
20.615528128088304
```

## Methods continued

You can declare a method on non-struct types, too.

In this example we see a numeric type `MyFloat` with an `Abs` method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

```go
package main

import (
   "fmt"
   "math"
)

type Myfloat float64

func (f Myfloat)Abs()float64  {
   if f < 0{
      return float64(-f)
   }
   return float64(f)
}

func main()  {
   f := Myfloat(-math.SqrtE)
   fmt.Println(f.Abs())
}
```

Output :

```
1.6487212707001282
```

## Pointer receivers

You can declare methods with pointer receivers.

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)

For example, the `Scale` method here is defined on `*Vertex`.

Methods with pointer receivers can modify the value to which the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the `*` from the declaration of the `Scale` function on line 16 and observe how the program's behavior changes.

With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behavior as for any other function argument.) The `Scale` method must have a pointer receiver to change the `Vertex` value declared in the `main` function.

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	 v.X = v.X * f
	 v.Y = v.Y * f
}

func main()  {
	v := Vertex{3, 4}
	v.Scale(104)
	fmt.Println(v.Abs())
}
```

Outpt :

```
520
```

## Pointers and functions

Here we see the `Abs` and `Scale` methods rewritten as functions.

Again, try removing the `*` from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

```go
package main

import (
   "fmt"
   "math"
)

type Vertex struct {
   X , Y float64
}

func Abs(v Vertex) float64{
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64)  {
   v.X = v.X * f
   v.Y = v.Y * f
}

func main() {
   v := Vertex{3,4}
   Scale(&v, 4)
   fmt.Println(Abs(v))
}
```

Output :

```
20
```

## Methods and pointer indirection

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

```
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

For the statement `v.Scale(5)`, even though `v` is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

```go
package main

import "fmt"

type Vertex struct {
   X ,Y float64
}

func (v *Vertex) Scale(f float64) {
   v.X = v.X * f
   v.Y = v.Y * f
}

func ScalFunc(v *Vertex, f float64) {
   v.X = v.X * f
   v.Y = v.Y * f
}

func main() {
   v := Vertex{3,4}
   v.Scale(4)
   ScalFunc(&v,4)

   fmt.Println(v)

   p := &Vertex{3,4}
   p.Scale(4)
   ScalFunc(p,4)
   fmt.Println(p)
   fmt.Println(*p)

}
```

Output :

```
{48 64}
&{48 64}
{48 64}
```

## Methods and pointer indirection (2)

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

```
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

while methods with value receivers take either a value or a pointer as the receiver when they are called:

```
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.

```go
package main

import "fmt"

type Vertex struct {
   X , Y float64
}

func (v *Vertex) Scale(f float64) {
   v.X = v.X * f
   v.Y = v.Y * f
}

func ScaleFunc(v Vertex, f float64) Vertex {
   v.X = v.X * f
   v.Y = v.Y * f
   return v
}

func main() {
   v := Vertex{3,4}
   v.Scale(4)
   fmt.Println(v)
   fmt.Println(ScaleFunc(v,4))

   p := &Vertex{3,4}
   p.Scale(4)
   fmt.Println(p)
   fmt.Println(ScaleFunc(*p,4))

}
```

Output :

```
{12 16}
{48 64}
&{12 16}
{48 64}
```

## Choosing a value or pointer receiver

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both `Scale` and `Abs` are with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

```go
package main

import (
   "fmt"
   "math"
)

type Vertex struct {
   X, Y float64
}

func (v *Vertex)Scale(f float64)  {
   v.X = v.X * f
   v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
   return  math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
   v := &Vertex{3,4}
   fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
   v.Scale(5)
   fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

Output :

```
Before scaling: &{X:3 Y:4}, Abs: 5
After scaling: &{X:15 Y:20}, Abs: 25
```

## Interfaces

An *interface type* is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

**Note:** There is an error in the example code on line 22. `Vertex` (the value type) doesn't implement `Abser` because the `Abs` method is defined only on `*Vertex` (the pointer type).

```go
package main

import (
   "fmt"
   "math"
)

type Abser interface{
   Abs() float64
}

func main(){
   var a Abser
   f := MyFloat(-math.Sqrt2)
   v := Vertex{3,4}
   fmt.Println(f)
   a = f  // a MyFloat implements Abser

   a = &v // a *Vertex implements Abser

   // In the flowing line, v is a Vertex(not *Vertex)
   // and does not implement Abser
   a = &v
   fmt.Println(a.Abs())

}



type MyFloat float64

func (f MyFloat) Abs() float64 {
   if f < 0{
      return float64(-f)
   }
   return float64(f)
}

type Vertex struct {
   X, Y float64
}

func (v *Vertex) Abs() float64  {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

Output :

```
-1.4142135623730951
5
```

## Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

```go
package main

import "fmt"

type I interface{
   M()
}

type T struct{
   S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so
func (t T) M()  {
   fmt.Println(t.S)
}

func main() {
   var i I = T{"mj"}
   i.M()
   i = T{"i love u"}
   i.M()
}
```

Output :

```
mj
i love u
```

## Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```
(value, type)
```

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

```go
package main

import (
   "fmt"
   "math"
)

type I interface {
   M()
}

type T struct {
   S string
}

func (t *T)M()  {
   fmt.Println(t.S)
}

type F float64

func (f F)M()  {
   fmt.Println(f*10)
}

func main() {
   var i I
   i = &T{"Hello"}
   describe(i)
   i.M()

   i = F(math.Pi)
   describe(i)
   i.M()
}

func describe(i I)  {
   fmt.Printf("(%v, %T)\n", i, i)
}
```

Output :

```
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
31.41592653589793
```

## Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method `M` in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.

```go
package main

import (
   "fmt"
)

type I interface {
   M()
}
type T struct {
   S string
}
func (t *T) M()  {
   if t == nil{
      fmt.Println("<nil>")
      return
   }
   fmt.Println(t.S)
}

func main() {
   var i I
   var t *T
   i = t
   describe(t)
   i.M()

   t = &T{"mj"}
   t.M()
}

func describe(i I)  {
   fmt.Printf("(%v, %T)\n",i, i)
}
```

Output :

```
(<nil>, *main.T)
<nil>
mj
```

## Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which *concrete* method to call.

```go
package main

import "fmt"

type I interface {
   M()
}

type T struct {
   S string
}

func (t *T) M()  {
   //if t == nil{
   // return
   //}
   fmt.Println(t.S)
}

func main() {
   var i I
   describe(i)
   var t *T
   describe(t)
   i = t
   i.M()
}

func describe(i I) {
   fmt.Printf("(%v, %T)", i ,i)
}
```

Output :

```
(<nil>, <nil>)(<nil>, *main.T)panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x109ea26]

goroutine 1 [running]:
main.(*T).M(0x0)
        /Users/henry/go/src/awesomeProject/nil-interface-values.go:25 +0x26
main.main()
        /Users/henry/go/src/awesomeProject/nil-interface-values.go:34 +0x123
```

## The empty interface

The interface type that specifies zero methods is known as the *empty interface*:

```
interface{}
```

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 2
	describe(i)

	i = "mj"
	describe(i)
}

func describe(i interface{})  {
	fmt.Printf("(%v, %T)\n",i ,i)
}
```

Output :

```
(<nil>, <nil>)
(2, int)
(mj, string)
```

## Type assertions

A *type assertion* provides access to an interface value's underlying concrete value.

```
t := i.(T)
```

This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic.

To *test* whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

```
t, ok := i.(T)
```

If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true.

If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

```go
package main

import (
   "fmt"
)

func main() {
   var i interface{} =  "mj"

   s := i.(string)
   fmt.Println(s)

   s, ok := i.(string)
   fmt.Println(s,ok)

   f, ok := i.(float64)
   fmt.Println(f, ok)

   f = i.(float64)  // panic
   fmt.Println(f)

}
```

Output :

```
mj
mj true
0 false
panic: interface conversion: interface {} is string, not float64

goroutine 1 [running]:
main.main()
```

## Type switches

A *type switch* is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.

This switch statement tests whether the interface value `i` holds a value of type `T` or `S`. In each of the `T` and `S` cases, the variable `v` will be of type `T` or `S` respectively and hold the value held by `i`. In the default case (where there is no match), the variable `v` is of the same interface type and value as `i`.

```go
package main

import "fmt"

func do(i interface{}) {

   switch v := i.(type) {
   case int:
      fmt.Printf("Third %v is %v\n", v, v*3)
   case string:
      fmt.Printf("%q is %v bytes long\n",v, len(v))
   default:
      fmt.Printf("I don't know about type %T\n", v)
   }
}

func main(){
   do("mj")
   do(520)
   do(true)
}
```

Output :

```
"mj" is 2 bytes long
Third 520 is 1560
I don't know about type bool
```

## Stringers

One of the most ubiquitous interfaces is [`Stringer`](https://golang.org/pkg/fmt/#Stringer) defined by the [`fmt`](https://golang.org/pkg/fmt/) package.

```
type Stringer interface {
    String() string
}
```

A `Stringer` is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values.

```go
package main

import "fmt"

type Person struct{
   Name string
   Age int
}

func (p Person)String() string{
   return fmt.Sprintf("%v (%v years)",p.Name, p.Age)
}

func main() {
   m := Person{"mj", 18}
   h := Person{"henry",25}
   fmt.Println(m,h)
}
```

Output :

```
mj (18 years) henry (25 years)
```

## Exercise: Stringers

Make the `IPAddr` type implement `fmt.Stringer` to print the address as a dotted quad.

For instance, `IPAddr{1, 2, 3, 4}` should print as `"1.2.3.4"`.

```go
package main

import (
   "fmt"
   "strconv"
   "strings"
)

type IPAddr [4]byte

//func (ip IPAddr) String() string {
// return fmt.Sprintf("%d.%d.%d.%d",ip[0],ip[1],ip[2],ip[3])
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
```

Output :

```
loopback: 127.0.0.1
googleDNS: 8.8.8.8
```

## Errors

Go programs express error state with `error` values.

The `error` type is a built-in interface similar to `fmt.Stringer`:

```
type error interface {
    Error() string
}
```

(As with `fmt.Stringer`, the `fmt` package looks for the `error` interface when printing values.)

Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.

```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

A nil `error` denotes success; a non-nil `error` denotes failure.

```go
package main

import (
   "fmt"
   "time"
)

type MyError struct{
   When time.Time
   What string
}

func (e *MyError) Error() string {
   return fmt.Sprintf("at %v, %s",e.When,e.What)
}

func run() error {
   return &MyError{
      time.Now(),
      "it didn't work!",
   }
}

func main() {
   if err := run(); err != nil{
      fmt.Println(err)
   }
}
```

Output :

```
at 2020-03-08 15:29:36.997841 +0800 CST m=+0.000085970, it didn't work!
```

## Exercise: Errors

Copy your `Sqrt` function from the [earlier exercise](https://tour.golang.org/flowcontrol/8) and modify it to return an `error` value.

`Sqrt` should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

```
type ErrNegativeSqrt float64
```

and make it an `error` by giving it a

```
func (e ErrNegativeSqrt) Error() string
```

method such that `ErrNegativeSqrt(-2).Error()` returns `"cannot Sqrt negative number: -2"`.

**Note:** A call to `fmt.Sprint(e)` inside the `Error` method will send the program into an infinite loop. You can avoid this by converting `e` first: `fmt.Sprint(float64(e))`. Why?

Change your `Sqrt` function to return an `ErrNegativeSqrt` value when given a negative number.

```go
package main

import (
   "fmt"
   "math"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
   return fmt.Sprintf("can not sqrt negative number %f", e)
}
const e = 1e-8
func Sqrt(x float64) (float64,error) {
   if x < 0{
      return 0,ErrNegativeSqrt(x)
   }

   z := x
   for{
      newZ := z - ((z*z) - x) / (z*2)
      if math.Abs(newZ- z) < e{
         return newZ,nil
      }
      z = newZ
   }

}

func main() {
   fmt.Println(Sqrt(2))
   fmt.Println(Sqrt(-2))
}
```

Output :

```
1.4142135623730951 <nil>
0 can not sqrt negative number -2.000000
```

## Readers

The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data.

The Go standard library contains [many implementations](https://golang.org/search?q=Read#Global) of these interfaces, including files, network connections, compressors, ciphers, and others.

The `io.Reader` interface has a `Read` method:

```
func (T) Read(b []byte) (n int, err error)
```

`Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.

The example code creates a [`strings.Reader`](https://golang.org/pkg/strings/#Reader) and consumes its output 8 bytes at a time.

```go
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
```

Output :

```
8
n = &{Hello Reader! 8 -1} err = <nil> b = [72 101 108 108 111 32 82 101]
b[:n] = "Hello Re"
5
n = &{Hello Reader! 13 -1} err = <nil> b = [97 100 101 114 33 32 82 101]
b[:n] = "ader!"
0
n = &{Hello Reader! 13 -1} err = EOF b = [97 100 101 114 33 32 82 101]
b[:n] = ""
```

## Exercise: Readers

Implement a `Reader` type that emits an infinite stream of the ASCII character `'A'`.

```go
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
```

Ouutput :

```
Ok!
```

## Exercise: rot13Reader

A common pattern is an [io.Reader](https://golang.org/pkg/io/#Reader) that wraps another `io.Reader`, modifying the stream in some way.

For example, the [gzip.NewReader](https://golang.org/pkg/compress/gzip/#NewReader) function takes an `io.Reader` (a stream of compressed data) and returns a `*gzip.Reader` that also implements `io.Reader` (a stream of the decompressed data).

Implement a `rot13Reader` that implements `io.Reader` and reads from an `io.Reader`, modifying the stream by applying the [rot13](https://en.wikipedia.org/wiki/ROT13) substitution cipher to all alphabetical characters.

The `rot13Reader` type is provided for you. Make it an `io.Reader` by implementing its `Read` method.

```go
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
```

Output :

```
URYYB, zw
```

