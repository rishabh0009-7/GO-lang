
//7- array in golang 

package main 
import "fmt"

func main (){
	var name [5] string //declaring array
	name[0]= "pamkaj"
	name[1]= "rahul"

	fmt.Println("names  of person is :",name)

	var numbers =[5] int{1,2,3,4,5} // initialising array 
	fmt.Println("numbers is ", numbers)

	// length of array 
	fmt.Println("length of array is ",len(numbers))

	//accesing array 
	fmt.Println("value of name at 1st  index is ", name[1])

	//agar humne kisi array ko koi value nhi di to uska deafult value 0 hota hai agar int diya aur string diya to "  " aur agar boolean diya to false 
	var price [5] int
	fmt.Println("given price is ",price)
}
