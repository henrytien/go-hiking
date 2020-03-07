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

