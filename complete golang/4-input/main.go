
//4- how to take input from user in golang

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var name string
	// fmt.Scan(&name)
	// fmt.Println("hello mr.", name)
	//but scan jo hai vo bas whitespace tak padta hai vo rishab srivastav nhi akr paayega print error dedega 
	// instead of above use bufio package 

	reader:=bufio.NewReader(os.Stdin)
	name, _:= reader.ReadString('\n')
	fmt.Println("hello mr",name)



}
