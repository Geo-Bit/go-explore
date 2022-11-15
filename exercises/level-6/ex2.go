// Create a func with identifer foo that takes in a variadic parameter of type int
// pass in a value of type []int into your func (unfurl the []int)
// Returns the sum of all vlaues of type int passed in
// Create a func with the identifier bar that take a parameter of type []int
// Returns the sum of all values of type int passed in

package main

import "fmt"

func main(){
	x := []int{1,2,3,4,5,6,7,}
	x2 := []int{1,2,6,4,5,6,7,}
	fmt.Println(foo(x...))
	fmt.Println(bar(x2))
}

func foo(xi ...int) int {
	sum := 0
	for _, v  := range xi{
		sum += v
	}
	return sum
}

func bar( i []int) int{
	sum := 0
	for _, v := range i{
		sum += v
	}
	return sum
}