// data conversion
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("numbeer is ", num)
	fmt.Printf("type of num is %T\n", num)
	// num= num +1.23 // nhi hoga error aayega isliye data conversion karna padega
	//data conversion- here we are converting int to float
	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("data is ", data)
	fmt.Printf("type of data is %T\n", data)

	//data conversion - here we are converting int to string or str to int
	//we use strconv.itoa function for int to string
	//use strconv.atoi for string to int
	//use strconv.parsefloat for string to float
// converting  int to string 
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("string  is ", str)
	fmt.Printf("type of string  is %T\n", str)
//converting string to int
	number_string :="1234"
	number_int ,_     := strconv.Atoi(number_string)
	fmt.Println("integer  is ", number_int)
	fmt.Printf("type of integer is %T\n", number_int)


	//converting string to float
	number_str:= "1234"
	number_float64,_:= strconv.ParseFloat(number_str,64)
	fmt.Println("float is ", number_float64)
	fmt.Printf("type of float64 is %T\n", number_float64)






}
