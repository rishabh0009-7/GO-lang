// maps similar to dictionary in python and hashmap in other programming language
// maps- unordered collection of key value pair
package main
import "fmt"


func main (){
	fmt.Println("maps")
	//how to make map
	//here we are mapping name<->grade

	studentgrades:= make(map[string]int)
	studentgrades["rishabh"]= 100
	studentgrades["rahul"]= 40
	studentgrades["rohan"]= 23
	fmt.Println("marks for rishabh is",studentgrades["rishabh"])

	//we can aslo update value 
	studentgrades["rahul"]= 54
	fmt.Println("marks for rahul is :", studentgrades["rahul"])
// we can also delete any key
// delete(studentgrades,"rahul")
// fmt.Println("marks for rahul is :", studentgrades["rahul"])

//checking if key exists
grades,exists :=studentgrades["david"]
fmt.Println("grades of david is ",grades)
fmt.Println("david exists",exists)



//loop ki madad se print karana ahi uppar ke key value pairs
for index, value := range studentgrades{
fmt.Printf("key is % s and marks is %d\n",index ,value)
}


//creating a map using a literal
 person:= map[string ]int {
	"alice": 90,
	"raam": 87,
	"rohan ": 76,
 }
 for index,value := range person{
	fmt.Printf("key is % s and marks is %d\n",index ,value)
 }
}