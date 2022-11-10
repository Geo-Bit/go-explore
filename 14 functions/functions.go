// Functions
// func (r receiver) identifier (parameters) (return(s)){...code...} <--- Syntax
// functions vs. parameters
// we define our func with parameters (if any)
// we call our func with arguments (if any)
// everything in Go is PASS BY VALUE

// package main

// import "fmt"

// Function Syntax
// func main() {
// 	foo()
// 	bar("Geo")
// 	s1 := woo("James Bond")
// 	fmt.Println(s1)
// 	x, y := mouse("James", "Bond")
// 	fmt.Println(x)
// 	fmt.Println(y)
// }

// func foo() {
// 	fmt.Println("This is foo.")
// }

// func bar(s string) {
// 	fmt.Println("Hello,", s)
// }

// func woo(s string) string {
// 	return fmt.Sprint("Hello from woo,", s)
// }

// func mouse(fn string, ln string) (string, bool) {
// 	a := fmt.Sprint(fn, " ", ln, " says 'Hello'")
// 	b := true
// 	return a, b
// }

// Variadic Parameters
// func main() {
// 	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
// 	fmt.Println("The total is", x)
// }

// func sum(x ...int) int { // ... is unlimited number of ints (variadic parameters)
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)

// 	sum := 0
// 	for i, v := range x {
// 		sum += v
// 		fmt.Println("for item in index position", i, "we are now adding", v, "to the total, which is now", sum)
// 	}
// 	return sum
// }

//Unfurling a Slice
// func main() {
// 	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
// 	fmt.Println("The total is", x)
// }

// func sum(x ...int) int { // ... is unlimited number of ints (variadic parameters)
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)

// 	sum := 0
// 	for i, v := range x {
// 		sum += v
// 		fmt.Println("for item in index position", i, "we are now adding", v, "to the total, which is now", sum)
// 	}
// 	return sum
// }

// Defer
// Defers the execution of a function until wherever its being called comes to an end
// EX: if you open a file in a program, you can immediately defer close it to ensure it gets closed
// at the end of the program
// package main

// import "fmt"

// func main() {
// 	defer foo() // deferred stuff gets run as soon as the function exits
// 	bar()
// }

// func foo() {
// 	fmt.Println("foo")
// }

// func bar() {
// 	fmt.Println("bar")
// }

// Methods
// package main

// import "fmt"

// // func (r receiver) identifier(parameters) return(s) {code...}
// func (s secretAgent) speak() { // attaches this function to secretAgent (receiver) structs
// 	fmt.Println("I am,", s.first, s.last)

// }

// type person struct {
// 	first string
// 	last  string
// }

// type secretAgent struct {
// 	person
// 	ltk bool
// }

// func main() {
// 	sa1 := secretAgent{
// 		person: person{
// 			"James",
// 			"Bond",
// 		},
// 		ltk: true,
// 	}
// 	sa2 := secretAgent{
// 		person: person{
// 			"Someone",
// 			"Else",
// 		},
// 		ltk: true,
// 	}

// 	fmt.Println(sa1)
// 	sa1.speak()
// 	sa2.speak()
// }

// Interfaces & Polymorphism
// Interfaces allow us to define behavior
// Interfaces allow us to do Polymorphism
// Interfaces - "Hey baby, if you have these methods, then you're my type"
// package main

// import "fmt"

// // func (r receiver) identifier(parameters) return(s) {code...}
// func (s secretAgent) speak() { // attaches this function to secretAgent (receiver) structs
// 	fmt.Println("I am,", s.first, s.last)

// }

// type person struct {
// 	first string
// 	last  string
// }

// type secretAgent struct {
// 	person
// 	ltk bool
// }

// // a value can be of more than one type
// type human interface { // Any type that has method speak() is also of type human
// 	speak() // since secret agent has a method speak(), then it is inherently human
// }

// func bar(h human) {
// 	fmt.Println(" I was passed into bar", h)
// }

// func (p person) speak() {
// 	fmt.Println("I am,", p.first, p.last, "- the person speak")
// }

// type hotdog int

// func main() {
// 	sa1 := secretAgent{
// 		person: person{
// 			"James",
// 			"Bond",
// 		},
// 		ltk: true,
// 	}
// 	sa2 := secretAgent{
// 		person: person{
// 			"Someone",
// 			"Else",
// 		},
// 		ltk: true,
// 	}

// 	p1 := person{
// 		first: "Dr.",
// 		last:  "No",
// 	}

// 	fmt.Println(sa1)
// 	sa1.speak()
// 	sa2.speak()

// 	fmt.Println(p1)
// 	bar(sa1) // poly (many) morph (change) this function can take in many types
// 	bar(sa2)
// 	bar(p1)

// 	//conversion
// 	var x hotdog = 42
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)
// 	var y int = int(x)
// 	fmt.Println(y)
// 	fmt.Printf("%T\n", y)

// }

// Anonymous Func
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Ayooo")
// 	foo() //not anonymous

// 	func(x int) { //anonymous
// 		fmt.Println("Hello from anonymous func", x)
// 	}(42)

// }

// func foo() { // not anonymous
// 	fmt.Println("Hello from foo")
// }

// Func Expression
// Assign a function to a variable <-- This
// Funcs are first class citizens (can be used like any other type or variable)
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Ayooo")
// 	f := func() {
// 		fmt.Println("my first func expression")
// 	}
// 	f()
// 	g := func(x int) {
// 		fmt.Println("The year big brother started watching: ", x)
// 	}
// 	g(1984)

// }

// Returning a Func
// Returning a function from a function
// package main

// import "fmt"

// func main() {
// 	x := bar()
// 	fmt.Printf("%T\n", bar())
// 	fmt.Println(x())
// 	fmt.Println(bar()()) //simplified
// 	// x := func() int { //anonymous
// 	// 	return 451
// 	// }
// 	//fmt.Printf("%T", x)

// }

// func foo() string {
// 	s := "Hello from foo" // unnecessary (just return the string)
// 	return s              // returning a string
// }

// func bar() func() int { // bar returns a function which returns an int
// 	return func() int {
// 		return 451
// 	}
// }

// Callback
// Passing a func as an argument
// Sometimes known as functional programming
// package main

// import "fmt"

// func main() {
// 	ii := []int{1, 2, 3, 4}
// 	s := sum(ii...) // must unfurl the slice of int
// 	fmt.Println("all numbers:", s)

// 	s2 := even(sum, ii...)
// 	fmt.Println("even numbers:", s2)

// 	s3 := odd(sum, ii...)
// 	fmt.Println("odd numbers:", s3)
// }

// func sum(xi ...int) int {
// 	fmt.Printf("%T\n", xi)
// 	total := 0
// 	for _, v := range xi {
// 		total += v
// 	}
// 	return total
// }

// func even(f func(xi ...int) int, vi ...int) int { // Callback
// 	var yi []int
// 	for _, v := range vi {
// 		if v%2 == 0 {
// 			yi = append(yi, v)
// 		}
// 	}
// 	return f(yi...)

// }

// func odd(f func(xi ...int) int, vi ...int) int { // Callback
// 	var yi []int
// 	for _, v := range vi {
// 		if v%2 != 0 {
// 			yi = append(yi, v)
// 		}
// 	}
// 	return f(yi...)

// }

// Closure
// Enclose the scope of a variable and contain it to a certain area
// Enclose some code around the variable to limit the scope of the variable
// package main

// import "fmt"

// // var x int // declare x, wider scope

// func main() {
// 	var x int      // narrower scope to func main()
// 	fmt.Println(x) // 0
// 	x++            // increment x here
// 	{              // limiting the scope of y here
// 		y := 1
// 		fmt.Println(y) // note can print y here since its in scope
// 	}
// 	// fmt.Println(y) // ^^ but cant print y here because out of scope
// 	fmt.Println(x) // 1
// 	foo()          // note the scope of x (it is also in foo)
// 	fmt.Println(x) // 2

// 	a := incrementor()
// 	b := incrementor()
// 	fmt.Println(a())
// 	fmt.Println(a())
// 	fmt.Println(b())
// 	fmt.Println(b())

// }

// func foo() {
// 	// x++ // increment x here
// 	fmt.Println("foo")
// }

// func incrementor() func() int { // incrementor example with closure
// 	var x int
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// Recursion
// Anything you do with recursion, you can do with loops
// Recursion is often times overly complex and can have memory drawbacks
// Loops > Recursion (usually)
// Factorial is the famous example
// package main

// import "fmt"

// func main() {
// 	fmt.Println(4 * 3 * 2 * 1) // Basic factorial (24)
// 	n := factorial(4)          // 24
// 	fmt.Println(n)
// 	n2 := factorialLoop(4) // loop instead
// 	fmt.Println(n2)        // 24
// }

// func factorial(n int) int {
// 	if n == 0 {
// 		return 1 // 4 * 3 * 2 * 1 * 1 <-- line 398
// 	}
// 	return n * factorial(n-1)
// }

// func factorialLoop(n int) int {
// 	total := 1
// 	for ; n > 0; n-- {
// 		total *= n
// 	}
// 	return total
// }
