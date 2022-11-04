// Create a slice of a slice of string. Store the following data in the multi-dimensional slice/
// "James", "Bond", "Shaken, not stirred"
// "Ok", "Great","Decent"
// Range over the records, then range over the data in each record
package main

import "fmt"

func main() {
	sliceOne := []string{"James", "Bond", "Shaken, not stirred"}
	sliceTwo := []string{"Ok", "Great", "Decent"}
	fmt.Println(sliceOne)
	fmt.Println(sliceTwo)
	sliceOneTwo := [][]string{sliceOne, sliceTwo}
	fmt.Println(sliceOneTwo)

	for i, slice := range sliceOneTwo {
		fmt.Println("record: ", i)
		for j, val := range slice {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

}
