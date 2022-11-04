// Struct is a data structure that allows you to combine different types of data AKA COMPOSITE Data structure
// Aggregate data type/structure
// Comparable to an object/class

package main

// Struct Introduction
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

// func main() {
// 	p1 := person{
// 		first: "James",
// 		last:  "Bond",
// 		age:   32,
// 	}
// 	p2 := person{
// 		first: "Someone",
// 		last:  "Else",
// 		age:   22,
// 	}

// 	fmt.Println(p1, p2)
// 	fmt.Println(p1.first, p2.first)

// }

// Embedded Structs (Embed a struct in another struct)
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

// type secretAgent struct { // embed a person in the secretAgent struct
// 	person
// 	licenseToKill bool
// }

// func main() {
// 	sa1 := secretAgent{
// 		person: person{
// 			first: "James",
// 			last:  "Bond",
// 			age:   32,
// 		},
// 		licenseToKill: true,
// 	}

// 	p2 := person{
// 		first: "Someone",
// 		last:  "Else",
// 		age:   22,
// 	}

// 	fmt.Println(sa1, p2)
// 	fmt.Println(sa1.first, p2.first)

// }

// Anonymous Structs (Include the struct definition within the main function instead of globally)
// For when you only need a struct to use once
// func main() {
// 	p1 := struct {
// 		first string
// 		last  string
// 		age   int
// 	}{
// 		first: "Someone",
// 		last:  "Else",
// 		age:   22,
// 	}

// 	fmt.Println(p1)

// }
