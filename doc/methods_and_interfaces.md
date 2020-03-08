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

