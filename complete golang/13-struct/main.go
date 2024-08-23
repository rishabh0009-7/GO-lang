
//structure is a composite datatype that groups together variables under a single name 
package main 
import "fmt"
 
type person struct{
	firstname string
	lastname string
	age int 
}


func main (){
var  rishabh person
rishabh.firstname= "rishabh"
rishabh.lastname = "srivastav"
rishabh.age =21
fmt.Println(" rishabh person:",rishabh)
//2nd method (shorter way) 

rahul:=person{
	firstname: "rahul",
	lastname: "agarwal",
	age: 23,
}
fmt.Println("rahul:",rahul)
// method 3 usingh new keyword
  var aman = new(person)
  aman.firstname = "aman "
  aman.lastname ="kumar"
  aman.age = 20 
  fmt.Println("aman:",aman)



  // you can also retrive any particular data 

  fmt.Println("aman age is :",aman.age)

}