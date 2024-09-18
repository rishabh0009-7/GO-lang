// crud- create ,read , update , delete
// api = jab humne app pe request kari vo server pe gyi fir server ne ek response bheja iss process ko  api karta hai
// we will use fake api kyuki yaha humne apna local server nhi bana rakha hai
// http.get
//get method- fetching information (viewing data / read data )
//ioutil.read

//crud api  - get method 

//method1
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main(){
	fmt.Println("crud api get method ")
	//1- ye pehle likha hai 
	response,err:= http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("erro getting",err)
		return

	}
	defer response.Body.Close() //ye last mai chalega 

// 3-code ka status check karna ke liye 
//agar humara link hi kharaab hai to ye dedega  
if response.StatusCode != http.StatusOK{
	fmt.Println("error is getting response :",response.Status)
	return

}



	// 2-ab jo response mila vo ioutil mai hai usko read karenge 

	data,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println("error reading",err)
		return
	}
	fmt.Println("data",string(data))
}