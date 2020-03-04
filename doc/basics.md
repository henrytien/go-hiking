

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

