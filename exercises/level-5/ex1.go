// Create your own type "person" which will have an underlying type of "struct" so that it can store:
// first name, last name, favorite ice cream flavors
// Create two values of TYPE person. Print out the values, ranging over the elements in the slice

package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

func main() {
	p1 := person{
		first:            "James",
		last:             "Bond",
		favoriteIceCream: []string{"vanilla", "chocolate"},
	}

	p2 := person{
		first:            "Someone",
		last:             "Else",
		favoriteIceCream: []string{"vanilla", "chocolate"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favoriteIceCream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favoriteIceCream {
		fmt.Println(i, v)
	}

}
