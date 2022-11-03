// Using a composition literal, create an array which holds 5 VALUES of TYPE int
package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	//x := make([]int, 5, 5)
	fmt.Println(x)
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	for i, v := range x {
		fmt.Println(i, v)
	}

}
