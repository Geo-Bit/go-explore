// Create a for loop using: for {} and have it print out the years you have been alive
package main

import "fmt"

func main() {
	yearOfBirth := 1992
	for {
		if yearOfBirth > 2022 {
			break
		} else {
			fmt.Println(yearOfBirth)
			yearOfBirth++
		}
	}
}
