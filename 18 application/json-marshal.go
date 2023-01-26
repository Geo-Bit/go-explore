package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Hulk",
		Last:  "Hogan",
		Age:   22,
	}

	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)

	// Marshal the data (convert to JSON)
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
