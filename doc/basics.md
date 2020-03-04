

## Packages

This program is using the packages with import paths `"fmt"` and `"math/rand"`.

**Note:** The environment in which these programs are executed is deterministic, so each time you run the example program `rand.Intn` will return the same number.

(To see a different number, seed the number generator; see [`rand.Seed`](https://golang.org/pkg/math/rand/#Seed). Time is constant in the playground, so you will need to use something else as the seed.)

## Imports

This code groups the imports into a parenthesized, "factored" import statement.

```go
import (
	"fmt"
	"math"
)
```

## Exported names	

In Go, a name is exported if it begins with a capital letter.

## Functions

A function can take zero or more arguments.

(For more about why types look the way they do, see the [article on Go's declaration syntax](https://blog.golang.org/gos-declaration-syntax).)

## Functions continued

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we shortened

```go
x int, y int
```

to

```go
x, y int
```

```go
package main

import "fmt"
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(4,9))
}
```

## Multiple results

A function can return any number of results. 

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main()  {
	//fmt.Println(swap("hello", "world"))
	a,b := swap("Hello", "World")
	fmt.Println(a, b)
}
```

## Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

```go
package main

import "fmt"
func split(sum int) (x, y int) {
	x = sum * 5 / 10
	y = sum - x
	return
}

func main() {
	fmt.Println(split(30))
}
```

## Variables

The `var` statement declares a list of variables; as in function argument lists, the type is last.

A `var` statement can be at package or function level. We see both in this example.

```go
package main

import (
	"fmt"
)

var Google, Apple ,Amazon bool
func main()  {
	var i int
	fmt.Println(i,Google, Apple, Amazon)
}
```

## Variables with initializers

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```go
package main

import "fmt"

var i, j int = 3, 4

func main() {
	var python, c, java  = true, false ,"yes"
	fmt.Println(i,j,python,c,java)
}
```

Output:

```go
3 4 true false yes
```

## Short variable declarations		

Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

```go
package main

import "fmt"

func main() {
	var i, j int = 3, 4
	c, python, cplusplus := "c", "python", "cplusplus"
	fmt.Println(i, j, c, python, cplusplus)
}
```

Output:

```go
3 4 c python cplusplus
```

## Basic types

Go's basic types are

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.

```go
package main

import (
	"fmt"
	"math/cmplx"
)
var	(
	Tobe bool = false
	MaxInt uint64 = 1 << 64-1
	z  complex128 = cmplx.Sqrt(-5 + 12i)
)
func main() {
	fmt.Printf("Type: %T Value: %v\n",Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n",MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n",z, z)
}
```

Output :

```go
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
```

## Zero values

Variables declared without an explicit initial value are given their *zero value*.

The zero value is:

`0` for numeric types,

`false` for the boolean type, and

`""` (the empty string) for strings.

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var mj bool
	var google string
	fmt.Printf("%v %v %v %q",i, f, mj, google)
}
```

Output :

```
0 0 false ""
```

## Type conversions

The expression `T(v)` converts the value `v` to the type `T`.

Some numeric conversions:

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, put more simply:

```
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the `float64` or `uint` conversions in the example and see what happens.

```go
package main

import (
   "fmt"
   "math"
)

func main() {
   var x, y int = 6, 9
   var f float64 = math.Sqrt(float64(x*x + y*y))
   var u uint = uint(f)
   fmt.Println(x, y, f, u)
}
```

Output :

```
6 9 10.816653826391969 10
```

## Type inference

When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

```
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant:

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

Try changing the initial value of `v` in the example code and observe how its type is affected.

```go
package main

import "fmt"

func main() {
   v := 42
   fmt.Printf("v is of type %T, value is of %v\n",v,v)
   f := 42.00
   fmt.Printf("f is of type %T, value is of %f\n",f,f)
}
```

Output :

```
v is of type int, value is of 42
f is of type float64, value is of 42.000000
```

## Constants

Constants are declared like variables, but with the `const` keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the `:=` syntax.

```go
package main

import "fmt"

const Pi = 3.14
func main() {
   const Word = "Google"
   fmt.Println("Hello",Word)
   fmt.Println("Happy", Pi, "Day")
}
```

Output :

```
Hello Google
Happy 3.14 Day
```

## Numeric Constants

Numeric constants are high-precision *values*.

An untyped constant takes the type needed by its context.

Try printing `needInt(Big)` too.

(An `int` can store at maximum a 64-bit integer, and sometimes less.)

```go
package main

import "fmt"

const(
   // Create a huger number by shifting a 1 bit left 100 places
   // In other words, the binary number is 1 followed by 100 zeros.
   Big = 1 << 100
   // Shift it right again 99 places, so we end with 1 << 1 or 2.
   Small = Big >> 99
)

func needInt(x int) int { return 10*x + 1}
func needFloat(x float64) float64 { return x*0.1}
func main() {
   fmt.Println(needInt(Small))
   fmt.Println(needFloat(Small))
   //fmt.Println(needInt(Big))  //overflows int
   fmt.Println(needFloat(Big))
}
```

Output :

```
21
0.2
1.2676506002282295e+29
```

## For

Go has only one looping construct, the `for` loop.

The basic `for` loop has three components separated by semicolons:

​		the init statement: executed before the first iteration

​		the condition expression: evaluated before every iteration

​		the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

The loop will stop iterating once the boolean condition evaluates to `false`.

**Note:** Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.

```go
package main

import "fmt"

func main() {
   sum := 0
   for i := 1; i <= 100; i++ {
      sum += i
   }
   fmt.Printf("Total is %v", sum)
}
```

Output :

```
Total is 5050
```

