// string packages - string functions 
package main

import (
	"fmt"
	"strings"
)


func main (){
	//splitting  the string using string function

	data := "apple, mango ,grapes, banana"
	parts:= strings.Split(data ,",")
	fmt.Println(parts)


	// if u want to count how many times word repeat in string 
	str:= "this is the word taht the same mssg"
	count := strings.Count(str,"the")
	fmt.Println("count",count)

	//to trim whitespace in a string

	str2:= "         hello       hi     /"
	trimm:= strings.TrimSpace(str2)
	fmt.Println("trimmed string is :",trimm)

	//if u want to join 2 strings  we use join but it uses array of string 
	firstname:= "rishabh"
	lastname:="srivastav"
	result:= strings.Join([]string{firstname,lastname} ," ")
	fmt.Println("result",result)

}