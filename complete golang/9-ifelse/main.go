//if-else condn

// package main

// import "fmt"

// func main() {
// 	x := 0
// 	if x > 5 {
// 		fmt.Println("x is greater than 5")

// 	} else {
// 		fmt.Println("x is less than 5")
// 	}
// }

//ifelseif

package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")

	} else if x > 2 {
		fmt.Println("x is greater than 2 but smaller than 5")
	} else {
		fmt.Println("x is  small")
	}

	y := 10
	if x > 5 && y > 8 { // we can also use operators
		fmt.Println("hello ")

	} else {
		fmt.Println("bye")
	}

}
