// Using ex2.go, use SLICING to create new slices test
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	fmt.Printf("%T\n", x) // print the type of x
	fmt.Println(x[:5])
	fmt.Println(x[2:6])
	fmt.Println(x[1:3])
	fmt.Println(x[:9])
}
