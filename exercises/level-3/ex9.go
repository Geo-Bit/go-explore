// Create a program that uses a switch statement with a switch expression specified as a variable of TYPE string with the IDENTIFIER "favSport"
package main

import "fmt"

func main() {
	favSport := "sports ball"
	switch favSport {
	case "basketball":
		fmt.Println("001")
	case "soccer":
		fmt.Println("002")
	case "sports ball":
		fmt.Println("003")
	}
}
