package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main(){
	fmt.Println("learning web request")

	//we use http.get for web request
	 res, err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	 if err!=nil{
		fmt.Println("error getting during get respnse ",err)
		return
	 }
	defer res.Body.Close()


	//error nhi ayaa tpo read karega 
	//reading the data 

	data,err:= ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("error ewhile reading",err)
		return
	}

	fmt.Println("response",string(data)) //data bytes ke form mai tha so we convert it into strings
}