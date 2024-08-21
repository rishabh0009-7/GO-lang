
//5- functions in golang

// jpo bhi hum function bahar banayega usko hume call  karna padega main function mai
package main 
import "fmt"

func simpleFunction(){
	fmt.Println("simple function")
}

func add(a,b int)(int ){ // 1st bracket is an input parameter and second one is an output type
	return a + b
}
//or
// func add(a int, b int )(result int ){

// 	result= a + b

// 	return

// }
//both are correct

func main(){
	simpleFunction()

	ans:= add(3,4)
	fmt.Println("add of 2 number is ", ans)
}