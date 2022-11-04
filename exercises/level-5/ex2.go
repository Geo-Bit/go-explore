// Take code from ex1 and store the values of type person in a map with the key of last name. Access each value in the map.
// Print out the values, ranging over the slice

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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favoriteIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("------------")
	}

}
