// Using a composition literal, create a slice of type int and assign 10 values. Then range over the slice and print the values out. Print out the type of the array
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x) // print the type of x
}
