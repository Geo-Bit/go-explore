package main

func main() {
	/* 	if true {
	   		fmt.Println("001")
	   	}
	   	if false {
	   		fmt.Println("002")
	   	}
	   	if !true {
	   		fmt.Println("003")
	   	}
	   	if !false {
	   		fmt.Println("004")
	   	}
	   	if 2 == 2 {
	   		fmt.Println("005")
	   	}
	   	if 2 != 2 {
	   		fmt.Println("006")
	   	}
	   	if !(2 == 2) {
	   		fmt.Println("007")
	   	}
	   	if !(2 != 2) {
	   		fmt.Println("008")
	   	} */

	// IF
	/* if x := 42; x == 2 { //; is an initialization statement
		fmt.Println("001")
	}
	//fmt.Println(x) doesn't work before of scope
	fmt.Println("002")
	fmt.Println("003") */

	// IF else
	/* 	x := 42
	   	if x == 40 {
	   		fmt.Println("001")
	   	} else {
	   		fmt.Println("002")
	   	} */

	// Else if
	/* 	x := 42
	   	if x == 40 {
	   		fmt.Println("001")
	   	} else if x == 42 {
	   		fmt.Println("002")
	   	} else {
	   		fmt.Println("003")
	   	} */

	// Modulus
	/* 	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	} */

	// Switch statements
	// No default fallthrough (below results in only 003)
	/* 	switch {
	   	case false:
	   		fmt.Println("001")
	   	case (2 == 4):
	   		fmt.Println("002")
	   	case (3 == 3):
	   		fmt.Println("003")
	   		fallthrough //enables fallthrough
	   	case (4 == 4):
	   		fmt.Println("004")
	   	default:
	   		fmt.Println("005")
	   	} */

	// Conditional logic operators
	/* 	fmt.Println(true && true)
	   	fmt.Println(true && false)
	   	fmt.Println(true || true)
	   	fmt.Println(true || false)
	   	fmt.Println(!true) */

}
