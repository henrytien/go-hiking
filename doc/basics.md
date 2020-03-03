

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

