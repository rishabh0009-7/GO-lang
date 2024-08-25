//pointer is a variable that stores address of other variable

package main
import "fmt"

//pointer ka function
func modifyvaluebyreference (num*int){
	*num= *num + 20

}

func main(){
	// fmt.Println("pointer")

	//creating a pointer 
	// var num int  //creating a variable
	// num:=2

	//method 1 
	// var ptr *int
	// ptr =&num

	//method 2
	// ptr := &num
//value se address pata lagaya 
	// fmt.Println("num has value :", num)
	// fmt.Println("pointer contains :",ptr)

	// // address se data ki value pata lagani hai 
	// fmt.Println("data contains in pointer :",*ptr)


	//if we only declare ptr but not assigned any vafriable in  it so it will store nil
	// var ptr *int// defaulgt pointer == nil 

	// creating a function for pointer 
	//modify value by reference means address se hi value ko change kardia    


		value :=10
		modifyvaluebyreference(&value) //fucntion call kiya hai uppar function banaya hai 
	fmt.Println("value contains",value)
}