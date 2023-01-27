// Pointers

// What are pointers?
// & gives the address
// * gives you the value stored at an address
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 42
// 	fmt.Println(a)
// 	fmt.Println(&a)        //print the memory address using & 0x140000140c8
// 	fmt.Printf("%T\n", a)  //int
// 	fmt.Printf("%T\n", &a) //*int - * means pointer to an int
// 	var b *int = &a
// 	fmt.Println(b)
// 	fmt.Println(*b)  // de-referencing an address using * and gives value in that address
// 	fmt.Println(*&a) // get address and de-reference to give value
// 	*b = 43
// 	fmt.Println(a)
// }

// When to use pointers
// Go is pass by value
package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

// Step 2. Pointer
func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}

// // Step 1. No pointer
// func foo(y int) {
// 	fmt.Println(y)
// 	y = 43
// 	fmt.Println(y)
// }

// Method Sets
// We can attach methods to a type (AKA a method set)
// Determine what methods attach to a TYPE
// a NON-POINTER RECEIVER works with values that are POINTERS or NON-POINTERS
// a POINTER RECEIVER only works with values that are POINTERS
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type circle struct {
// 	radius float64
// }

// type shape interface {
// 	area() float64
// }

// type (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func info(s shape) {
// 	fmt.Println("area",s.area())
// }

// func main() {
// 	c := circle{
// 		radius: 5,
// 	}
// 	info(c)
// }
