// 2-variables in go-lang
package main

import (
	"fmt"
)

func main() {
	var name string = "rishabh"
	var  currency = 6788 //datatype likh bhi skta hai aur nhi bhi faraq nhi padega
	fmt.Println(name)
	fmt.Println(currency)

	//we can aslo write const
	const naam ="harsh"
	fmt.Println(naam)

	//shortcut to write variable
	person:=123
	fmt.Println(person)

	//note- agar koi variable ko hume export karna hai yaa fir kisi aur package mai use karna hai to uska first letter capital mai likho
	//agar kis variable ko humne small letter se start kiya to usko hum bas ussi file mai acces kar skta hai bahar nhi
	var Public ="data is imp"
	fmt.Println(Public)

}