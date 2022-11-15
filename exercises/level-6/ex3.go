// Use the defer keyword to show that a deferred func runs after the func containing it exists

package main

import "fmt"

func main(){
	defer foo()
	fmt.Println("Yellow")
}

func foo(){
	defer func(){
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("Foo Ran")
}
