// Start with slice, append to that slice the value 52, print out the slice,
// in one statement append to that slice the values 53,54,55,
// Print out the slice, append to the slice the second slice, and print out slice
package main

import "fmt"

func main() {
	firstSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	firstSlice = append(firstSlice, 52)
	fmt.Println(firstSlice)
	firstSlice = append(firstSlice, 53, 54, 55)
	fmt.Println(firstSlice)
	secondSlice := []int{56, 57, 58, 59, 60}
	firstSlice = append(firstSlice, secondSlice...)
	fmt.Println(firstSlice)
}
