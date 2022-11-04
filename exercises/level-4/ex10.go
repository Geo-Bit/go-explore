// Using code from ex9, delete a record from the map
// Print using range
package main

import "fmt"

func main() {
	m := map[string][]string{
		"john doe": []string{"hunting", "fishing", "sportsball"},
		"suzie q":  []string{"hunting", "fishing", "sportsball"},
	}
	fmt.Println(m)

	m["Dr. No"] = []string{"Boats", "Skis", "Bees Knees"}

	delete(m, "Dr. No")

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
