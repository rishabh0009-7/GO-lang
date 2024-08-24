
//structure is a composite datatype that groups together variables under a single name 
package main 
import "fmt"
 
type person struct{
	firstname string
	lastname string
	age int 
}

type contact struct {
  email string
  phone  string

}

type address struct{
  house  int 
  area string
  state string
}

type employee struct{ // sabse bada struct jisma saare struct aa jayega
    person_details  person
    person_contact contact
    person_address address

}




func main (){
// var  rishabh person
// rishabh.firstname= "rishabh"
// rishabh.lastname = "srivastav"
// rishabh.age =21
// fmt.Println(" rishabh person:",rishabh)
//2nd method (shorter way) 

// rahul:=person{
// 	firstname: "rahul",
// 	lastname: "agarwal",
// 	age: 23,
// }
// // fmt.Println("rahul:",rahul)
// // method 3 usingh new keyword
//   var aman = new(person)
//   aman.firstname = "aman "
//   aman.lastname ="kumar"
//   aman.age = 20 
  // fmt.Println("aman:",aman)



  // you can also retrive any particular data 

  // fmt.Println("aman age is :",aman.age)


  //structure ke  andar structure bhi ghusa skta hai jaise employee ek bada structure banaya usma humne person, contact ,address daal diya hai 
   var employee1 employee
   employee1.person_details= person{
      firstname :"rishabh",
      lastname : "srivastav",
      age: 21,

      employee1.person_contact = contact{
        email: "srivastavrishab986@gmail.com",
        phone: "8287096769",
    }

    employee1.person_address= address{
      house: 12 ,
      area: "delhi",
       state:"delhi",
 
    }
  }
    fmt.Println("employee1 details:",employee1.)
   
  

   


}

