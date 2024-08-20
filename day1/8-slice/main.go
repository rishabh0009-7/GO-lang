//8- slice in golang (better way of array )
// isme hum array mai size nhi btata  aur array update kar skta hai normal array mai ye karna possible nhi tha 
package main 
import "fmt"

func main (){

	// numbers:=[]int {1,2,3,4,5,} //slice
	// numbers = append (numbers,6,7,8,9,0)
	// fmt.Println("number :",numbers)

	// method 2 using = make function 
	numbers:=make([] int ,3,5) // 3 is the length and 5 is the capacity 
	fmt.Println("slice", numbers)
	fmt.Println("length", len(numbers))
	fmt.Println("capacity", cap(numbers))

	// agar hum  kuch aur cheez add karta hai 5  capacity bharna ke baad to isme eerror nhi aayega capacity apna aap bada dega 10 kar dega
	// you cant initialise an empty slice 
	// var stock= []string //this is wrong always used make function
	// stock:=[]string{}//aise kar skta ho koi dikkat nhi
	
}