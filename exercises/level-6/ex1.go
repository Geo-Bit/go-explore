// Create a func with the identifier foo that returns an int
// Create a func with the identifer bar that returns an int and a string
// call both functions
// print out their results

package main

import "fmt"

func main(){
	fmt.Println(foo())
	fmt.Println(bar())

}

func foo() int{
	return 10
}

func bar() (int, string){
	return 1984, "Big Brother is Watching"
}