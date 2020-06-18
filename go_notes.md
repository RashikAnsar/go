# Go

[Golang](https://golang.org)
[Golang tour](https://tour.golang.org/)
[Digital Ocean](https://www.digitalocean.com/community/users/gopherguides)

## Creators

**Ken Thompson**, **Rob Pike** and **Robert Griesemar**.

## Features

- Strong (type of a variable cannot change over time) and statically (variables have to be defined at compile time) typed
- Simplicity
- Fast compile times
- Garbage collected
- Built-in Concurrency
- Compile to a standalone binaries

## Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Gopher!")
}
```

- `package main`: Every go file is gonna have that what package its a part of. Every programs entry point is main package
- `import "fmt"`: importing 'fmt' package
- `main()`: main function is entry point of the program

## Variables

### Variable declaration

- In GO, we must use all the declared variables or else compiler throws an error at compile time. (This keeps the go programs nice and clean)
- Variable can be declared in the following ways
  - `var <VARIABLE_NAME> type`
  - `var <VARIABLE_NAME> = value`
  - `<VARIABLE_NAME> := value` (short-hand notation invalid in package scope [only valid within the function blocks])
- In Go, all the declared variables will have the zeroed value of the given type.

> In second method there is no need to explicitly specify the type since, we are intializing the variable with value. GO infers the type of a variable from the initialized value.

```go
package main

import "fmt"

// package level variable
// it must have var keyword
var i float32 = 42

func main() {
    var year int;
    year = 2020

    var lang = "golang"

    version := 1.14

    fmt.Println(year, lang, version)

}
```

- For declaring multiple variables `var` keyword clutters the code so we can group them and declare

```go
// multiple variables are declared and using only one var keyword
var (
    actorName string = "Rober Downey jr."
    knownAs string = "Iron Man"
    sequels int = 3
    featuredIn string = "Marvel Cinematic Universe"
)
```

### Redeclaration and shadowing

- Redeclaration of a variable(within the same scope) throws a compile time error by compiler. [Can't redeclare a variable but we can shadow them]

```go
// shadowing example
package main

import "fmt"

// Package scope variable
var i int = 21

func main() {
    fmt.Println(i) // 21
    // function scope variable
    var i int = 42
    fmt.Println(i) // 42
}
```

### Visibility

- Naming of the identifier (variable (or) function name) controls the visibility in GO.
- If and only if the identifier name starts with capital letter they are exported from the package else it's visible only to that package (lowercase firstletter of variables are scoped to the package).

### Naming conventions

- Pascal or CamelCase

```go
var langName string = "Go"
var javaScript string = "ES7"
```

- Keep acronyms all upper case

```go
var theURL string = "https://golang.org"
```

- Name identifers in such a way that they are self explanatory

### Type conversions

- Converting data from one type to another.
- In GO, we've to do explicit conversion when changing type since, GO doesn't want to risk the losing of information while conversion.

```go
package main

import "fmt"

func main() {
    var i int = 42
    fmt.Printf("%v, %T\n", i, i)

    var j float32
    j = float32(i)
    fmt.Printf("%v, %T\n", j,j)
}
```

> In GO, string is a stream of bytes (unicode). To proper conversion we use `'strconv'` package

## Primitives

### Boolean Type

```go
package main

import "fmt"

func main() {
    var n bool = true
    fmt.Printf("%v, %T\n",n,n)
}
```

### Numeric Type

#### Integers

1. Signed Integers (`int`, `int8`, `int16`, `int32`, `int64`)
2. UnSigned Integers (`uint`, `uint8`, `uint16`, `uint32`)

- Arithmetic Operations: `+`, `-`, `*`, `/`, `%`
- Bitwise operators: `&`, `|`, `^`, `&^`, `<<`,`>>`

#### Floating Point numbers

- Float32 and Float64
- Literal styles:
  - Decimal (`3.14`)
  - Exponential (`13e18` or `2E10`)
  - Mixed (`13.7e12`)

#### Complex Numbers

- Complex64 and Complex128
- Built-in functions
  - `complex` - Make complex number from two floats
  - `real` - get real part as float
  - `imag` - get imaginart part as float

### Text Types

- String
  - UTF-8
  - Immutable
  - can be concatenated with (`+`) operator
  - can be converted to `[]byte` (slice of bytes)
- Rune
  - UTF-32
  - Alias for int32
  - Special methods normally require to process
    - eg: `strings.Reader#ReadRune` from Golang API

## Constants

- Constants cannot be reassigned. (immutable but can be shadowed)
- Constants has to be assignable at compile time.

### Naming Conventions

- Constants are named as similar to the variables. (PascalCase for exported constants, camelCase for internal constants)

```go
package main

import "fmt"

func main() {
    const myName = "Go"
}
```

### Typed Constants

- Work like immutable variables
- Can interoperate only with the same type.

```go
const myConst int8 = 42
```

### UnTyped Constants

- Work like literals
- Can interoperate with similar types

### Enumerated Constants

- Special symbol `iota` allows related constants to be created easily
- Iota starts at 0 in each const block and increment by one
- Watch out of constant values that match zero values for variables

```go
package main

import "fmt"

// const (
//     a = iota
//     b = iota
//     c = iota
// )

// The above const can also be written as (most common)
// the iota is scoped to constant block
const (
    a = iota
    b
    c
)

func main() {
    // iota changes its value as constants are evaluated
    fmt.Printf("%v\t %T\n", a, a)
    fmt.Printf("%v\t %T\n", b, b)
    fmt.Printf("%v\t %T\n", c, c)
}
```

#### Enumerated expressions

- Operations that can be determined at compile time are allowed
  - Arithmetic
  - Bitwise Operations
  - BitShifting etc.

```go
package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
    // dont forget . at the end
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}
```

## Arrays and Slices

### Arrays

- Collection of items of same type
- Fixed size
- Access via zero based index

#### Creation

```go
package main

import "fmt"

func main() {
	// grades := [3]int{97, 84, 69}
    grades := [...]int{97, 84, 69}
	fmt.Printf("Grades: %v\n", grades)
    var students [3]string
    fmt.Printf("Students: %v\n", students)
    students[0] = "Lisa"
    students[2] = "Andrew"
    students[1] = "Gopher"
    fmt.Printf("Students: %v\n", students)

    fmt.Printf("Number of students: %v\n", len(students))
}
```

```go
package main

import "fmt"

func main() {
	// var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix)
}
```

#### Built-in functions

- `len` function returns size of the array

#### Working with arrays

- Copies refer to different underlying data

### Slices

- Backed by array

#### Creation

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   //slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice of first 6 elements
	e := a[3:6] //slice of 4th, 5th and 6th element
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
    fmt.Println(e)

    x := make([]int, 10) //create a slice with capacity and length == 10
    y := make([]int, 10, 100) // slice with length == 10 and capacity == 100
}
```

#### Built-in functions

- `len` function returns length of slice
- `cap` function returns length of underlying array
- `append` function to add elements to slice (may cause expensive copy operation if underlying array is too small)

#### Working with slices

- Copies refer to the same underlying array

## Maps and Struct

### Maps

- Collections of value types that are accessed via keys
- Created via literals or via make function
- Members accessed via `[key]` syntax
  - `myMap['key'] = 'value'`
- Check for presence with `value, ok` from the result
- Multiple assignments refer to same underlying data

#### Creation

```go
package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(statePopulations)
}
```

### Structs

- Collection of disparate data types that describe a single concept
- Keyed by named fields
- Normally created as types, but anonymous structs are allowed
- Structs are value types
- No inheritance, but can use composition via embedding
- Tags can be added to struct fields to describe field

#### Creation

```go
package main

import "fmt"

// Doctor struct
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	aDoctor2 := Doctor{
		4,
		"Liz Shaw",
		[]string{
			"Jon Pertwee",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor2)
}
```

#### Embedding

```go
type User struct {
	name  string `example:"name"`
	email string `example:"example@mail.com"`
}

type Admin struct {
	User
	isAdmin bool
}
```

#### Tags

```go
type User struct {
    Name string `example:"name"`
}
```

## Control flow

### If-statements

```go
package main

import "fmt"

func main() {
	number := 50
	guess := -1

	if guess < 1 {
		fmt.Println("Guess must be greater than 1")
	} else if guess > 100 {
		fmt.Println("Guess must be less than 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else if guess == number {
			fmt.Println("You got it!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
}
```

### Switch statements

```go
package main

import "fmt"

func main() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
	case i >= 20:
		fmt.Println("greater than or equal to 20")
	default:
		fmt.Println("greater than twenty")
	}
}
```

### Looping

- **For loop**:
  - `for initializer; test; increment {}`
  - `for test {}`
  - `for {}`
- Exiting loop early:
  - break
  - continue
  - labels
- Looping over collections
  - arrays, slices, maps, strings, channels
  - `for k, v := range collection {}`

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
```

### Defer, Panic, Recover

#### Defer

- Defer follows **LIFO** order
- Defer used to delay execution of statement until function exits
- Useful to group **open** and **close** functions together
  - Be careful in loops
- Arguments evaluated at time defer is executed, not at the time of called function execution.

```go
package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}
```

```go
package main

import "fmt"

func main() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
```

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
```

#### Panic

- Occur when program cannot continue at all
  - Don't use when file can't be opened, unless it is critical
  - Use for unrecoverable events - cannot obtain TCP port for web server
- Function will stop executing
  - Deferred functions will still fire
- If nothing handles panic, program will exit

```go
package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
```

```go
package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	panic("Something bad happened")
	fmt.Println("end")
}
```

#### Recover

- Used to recover from panics
- Only useful in deferred functions
- Current function will not attempt to continue, but higher function in call stack will

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("[ERROR]:::", err)
		}
	}()
	panic("Something bad happend")
	fmt.Println("end")
}
```

## Pointers

### Creating

- Pointer types use an asterisk (`*`) as a prefix to typed pointed to
  - `*int` a pointer to an integer
  - Use the addressof operator (`&`) to get address of a variable

```go
package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	*b = 16
	fmt.Println(a, *b)
	fmt.Println(&a, b)

}
```

### Dereferencing

- Dereference a pointer by preceding with an asterisk (`*`)
- Complex types (e.g. structs) are automatically dereferenced

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
}
```

### The `new` function

- Create pointers to objects
  - Can use addressof operator (&) if value type already exists
    - `ms := myStruct{foo: 42}`
    - `p := &ms`
  - Use addressof operator before initializer
    - &myStruct{foo: 42}
  - Use the `new` keyword
    - Can't initialize the fields at the same time

```go
package main

import "fmt"

func main() {
	var ms *myStruct
	// ms = &myStruct{foo: 42}
	ms = new(myStruct)
	fmt.Println(ms)
	// (*ms).foo = 42
	// fmt.Println((*ms).foo)
	ms.foo = 42
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
```

### Working with nil

```go
package main

import "fmt"

func main() {
	var ms *myStruct
	fmt.Println(ms)
	// ms = &myStruct{foo: 42}
	ms = new(myStruct)
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
```

### Types with internal pointers

- All assignment operations in Go are copy operations
- Slices and maps contain internal pointers, so copies point to same underlying data

## Functions

### Basic Syntax

```go
package main

import "fmt"

func main() {
	greet()
}

func greet() {
	fmt.Println("Hello, World!")
}
```

### Parameters

```go
package main

import "fmt"

func main() {
	greeting, name := "Hello", "Go"
	greet("Hello, World!")
	// pass by value
	sayGreeting(greeting, name)
	// pass by reference
	sayGreetings(&greeting, &name)
	fmt.Println(name)

	sum("The sum is ", 1, 2, 3, 4, 5)

	total := sum2(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("The sum is ", total)
}

func greet(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting string, name string) {
	fmt.Println(greeting, name)
}

func sayGreetings(greeting *string, name *string) {
	fmt.Println(*greeting, *name)
	*name = "gopher"
	fmt.Println(*name)
}

func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}
```

### Return values

```go
package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("The sum is ", total)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
```

```go
func sum(values ...int) *int {
	result := 0

	for _, v := range values {
		result += v
	}

	return &result
}
```

```go
func sum(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}
```

```go
package main

import "fmt"

func main() {
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Prinln(err)
		return
	}
	fmt.Println(d)
}

func divide(a float64, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
```

### Anonymous Functions

```go
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello from anonymous func")
	}()

	a := func() {
		fmt.Println("From a")
	}
	a()
}
```

### Functions as types

```go
package main

import "fmt"

func main() {
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		} else {
			return a / b, nil
		}
	}

	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
```

### Methods

```go
package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
```

```go
package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
```

## Interfaces

```go
package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

}

// Writer interface
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter struct
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
```

```go
package main

import "fmt"

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

// Incrementer interface
type Incrementer interface {
	Increment() int
}

// IntCounter type
type IntCounter int

// Increment function
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
```

### Best Practices

- Use many, small interfaces
  - Single method interfaces are some of the most poweful and flexible
    - `io.Writer`, `io.Reader`, `interface{}`
- Dont export interfaces for types that will be consumed
- Do export interfaces for types that will be used by package
- Design functions and methods to recieve interfaces whenever possible

## Go Routines

### Creating

- Use `go` keyword infront of function call
- When using anonymous functions, pass data as local variables

### Synchronization

- Use `sync.WaitGroup{}` to wait for groups of goroutines to complete
- Use `sync.Mutex` and `sync.RWMutex` to protect data access

### Parallelism

- By default, Go will use CPU threads equal to available cores
- Change with `runtime.GOMAXPROCS`
- More threads can increase performance, but too many can slow it down

### Best Practices

- Dont create goroutines in libraries
  - Let consumer control concurrency
- When creating a goroutine, know how it will end
  - Avoid subtle memory leaks
- Check for race conditions at compile time

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
}
```

```go
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "GoodBye"
	wg.Wait()
}
```

```go
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
```

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	// m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
```

## Channels

### Channel Basics

- Create a channel with `make` command
  - `make(chan int)`
- Send message into channel
  - `ch <- val`
- Recieve message from channel
  - `val := <-ch`
- Can have multiple senders and recievers

### Restricting data flow

- Channel can be cast into send-only or receive only versions
  - Send-only: `chan <- int`
  - Receive-only: `<- chan int`

### Buffered channel

- Channels block sender side till receiver is available
- Block receiver side till message is available
- Can decouple sender and receiver with buffered channels
  - `make(chan int, 50)`
- Use buffered channels when send and receive have assymmetric loading

### For...range Loops with channels

- Use to monitor channel and process messages as they arrive

### Select statements

- Allow goroutines to monitor several channels at once
  - Blocks if all channels block
  - If multiple channels receive value simultaneously, behaviour is undefined.
