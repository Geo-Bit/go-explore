//Create a person struct
//Create a func called changeMe with a *person as a parameter
//Change the value stored at the *person address
//Create a value of type person
//Print out the value
//In func main
//Call changeMe
//In func main
//Print out the value
//Important
//to dereference a struct, use (*value).field
//p1.first
//(*p1).first
//“As an exception, if the type of x is a named pointer type and (*x).f
//is a valid selector expression denoting a field (but not a method),
//x.f is shorthand for (*x).f.”

package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}

func changeMe(p *person) {
	p.name = "Miss Moneypenny"
	(*p).name = "Miss Moneyp"
}
