//in method 1 jo hume json mil raha tha vo direct print hogya par iss method mai humne json ko  todo variable mai daala hai fir print kara
//http.get
//struct  mai new decoder   .decode
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//method 2 DECODE METHOD

type Todo struct { //startimg letter should be capital in structure varna error aayega 
	UserID   int    `json:"userId"`
	ID       int    `json:"id"`
	Title     string `json:"Title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("instead of  ioutil method 1 we will use method 2  using structure ")
	//1- ye pehle likha hai
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("erro getting", err)
		return

	}
	defer response.Body.Close() //ye last mai chalega

	// 3-code ka status check karna ke liye
	//agar humara link hi kharaab hai to ye dedega
	if response.StatusCode != http.StatusOK {
		fmt.Println("error is getting response :", response.Status)
		return

	}
	//2-using structure
	//new decoder -   kya karta hai ki jo bhi data rahega usko decode  karega normal objects mai aur usko save karega todo structure mai

	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo)

	if err != nil {
		fmt.Println("error decoding", err)
		return
	}

	fmt.Println("Todo:", todo)
	fmt.Println("title response:",todo.Title)
	fmt.Println("completed response",todo.Completed)

}
