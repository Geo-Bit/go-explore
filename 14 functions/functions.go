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
