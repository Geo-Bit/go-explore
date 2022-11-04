// To DELETE from a slice, we use APPEND along with SLICING.
// Start with the slice
// Use APPEND & SLICING to get the values to print out [42,43,44,48,49,50,51]
package main

import "fmt"

func main() {
	firstSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	firstSlice = append(firstSlice[:3], firstSlice[6:]...)
	fmt.Println(firstSlice)
}
