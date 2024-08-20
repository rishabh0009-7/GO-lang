
// 6- error handling + underscore identifier in golang (blank identifier )
// error handling stands for error ko handle karna hum apna hisaab se decide kar  skta hai ki hume kya error chahiye 


package main 
import "fmt"

func divide (a,b float64)(float64, error){ //  ye jo error hai iski jagah kuch aur bhi likh skta hai farak nhi padega 
	if b==0{
		return 0,fmt.Errorf("denominator  must not be zero ")
	}

	return a / b , nil //matlab ki agar den 0 nhi hai error to nil aayega na 
}
// or

// func divide (a,b float64)(float64, string ){ 
// 	if b==0{
// 		return 0,"denominator  must not be zero "
// 	}

// 	return a / b , nil //matlab ki agar den 0 nhi hai error to nil aayega na 
// }
// both are corect 


func main (){

	ans,_:= divide (10,0)
	fmt.Println("division of  2 number is ", ans)

	// underscore identifier jo hai vo variable ki tarah kaam karta hai jisme faltu cheezein daal skta hai  
// 2 output mai se agar ek kaa kaam nhi hai to usse hum underscore mai daal skta hai 

}

