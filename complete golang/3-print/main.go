// 3- difference between printf and println in golang
// print ln is used to print the text jaise hai parprint f is used to print in a formatted manner
package main
import
	"fmt"

func main(){
	age:=25
	name:= "rishabh"
	height := 5.2323
	fmt.Println("age :",age ,"height:",height,"name:",name)
	//age:25height:5.2323name:alice

	fmt.Printf("age is %d\n",age)//%d is for int kts a format specifier
	fmt.Printf("height is %.2f\n",height)//%.2f is for ans upto 2 decimal places
	fmt.Printf("type of height is %T\n",height)

	//another method for writing printf statements

	fmt.Printf("age: %d, height: %f, name: %s\n",age,height , name)

}