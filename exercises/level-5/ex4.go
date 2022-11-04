// Create and use an anonymous struct

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

func main() {
	sedan := struct {
		vehicle
		luxury bool
	}{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(sedan.luxury)

}
