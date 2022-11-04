// Create a new type vehicle with fields doors and color
// Create two additional types, truck & sedan
// Embed vehicle type in both truck & sedan
// Give truck the field fourWheel and sedan Luxury, both set to bool
// Using the vehicle, truck and sedan structs create a value of type truck and assign values
// create a value of type sedan and assign values to the fields
// Print out each of the values
// Print out a single field from each of the values

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: false,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(t1, s1)
	fmt.Println(t1.fourWheel)
	fmt.Println(s1.luxury)

}
