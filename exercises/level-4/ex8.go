// Create a map with a key of TYPE string which is a person's "last_first" name,
// and a value of TYPE []string which stores their favorite things. Store three
// records in your map. Print out all the values along with their index position
// in the slice.
package main

import "fmt"

func main() {
	m := map[string][]string{
		"john doe": []string{"hunting", "fishing", "sportsball"},
		"suzie q":  []string{"hunting", "fishing", "sportsball"},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
