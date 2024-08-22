// in golang there is only one loop i.e for loop
// no while loop // no do- while loop
// for loop mai hi while loop laga skta hai hum
package main

import "fmt"

func main() {
	//method 1
	//here we are printing 1 to 10
	for i := 0; i < 10; i++ {
		fmt.Println("numbers are :", i)
	}
	//method 2
	//here we  are printing infinite loop
	// counter:=0
	// for {
	// 	fmt.Println("infinite loop")
	// 	counter ++

	// }
	// infinite loop with break statement
	counter := 0
	for {
		fmt.Println("infinite loop")
		counter++
		if counter == 3 {
			break
		}

	}
	//range keyword print karne ke liye hota hai index and value
	//range keyword  - The range keyword is used in for loop to iterate over items of an array, slice,string ,  channel or map. With array and slices, it returns the index of the item as integer. With maps, it returns the key of the next key-value pair.
	numbers := []int{1, 22, 3, 44, 5}
	for index, value := range numbers {
		fmt.Printf("index of number is %d , and value is %d\n", index, value)
	}

	data := "hello world "
	for index, value := range data {
		fmt.Printf("index of data is %d  and value is %c\n", index, value)
	}
}
