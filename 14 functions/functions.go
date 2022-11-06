// Functions
// func (r receiver) identifier (parameters) (return(s)){...code...} <--- Syntax
// functions vs. parameters
// we define our func with parameters (if any)
// we call our func with arguments (if any)
// everything in Go is PASS BY VALUE

package main

import "fmt"

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
func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is", x)
}

func sum(x ...int) int { // ... is unlimited number of ints (variadic parameters)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total, which is now", sum)
	}
	return sum
}
