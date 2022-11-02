// Create a for loop using: "for condition {}" and have it print out the years you have been alive
package main

import "fmt"

func main() {
	yearOfBirth := 1992
	for yearOfBirth <= 2022 {
		fmt.Println(yearOfBirth)
		yearOfBirth++
	}
}
