package main 
import "fmt" 

// jiske aage defer keyword laga dia vo sabse end mai run hoga 

func main (){

	fmt.Println("starting of the program")
	fmt.Println("middle of the program")
	fmt.Println("end of the program")

	//output
	// starting of the program
// middle of the program
// end of the program

//using defer
fmt.Println("starting of the program")
	 defer fmt.Println("middle of the program")
	fmt.Println("end of the program")

	//output
	//starting of the program
// end of the program
// middle of the program

//agar 2 defer agya 
defer fmt.Println("middle of the program")
	defer fmt.Println("end of the program")


	//to ye stack ki tarah hojayega  aur jo sabse last mai defer hoga vo sabsse pehle aayega 





	 
}