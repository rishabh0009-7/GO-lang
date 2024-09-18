// handle url means  ek noermal stringh yani link ko hum packages ki madad se url mai convert krdenge
//example- url:="https://jsonplaceholder.typicode.com/todos/1" ye cheez hume  pata hai url hai par computer ke liye syring hai na

package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning url")
	myURL := "https://www.notion.so/Getting-Started-44204b9d92ca4f6ba08431c646ebcd50?cookie_sync_completed=true"
	fmt.Printf("type of this is %T\n", myURL)
	//to convert this string into url we use url.Parse
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("can't parse url", err)
		return
	}
	fmt.Printf("type of this is :%T\n", parsedURL)


//accesing url components

fmt.Println("scheme:", parsedURL.Scheme)
fmt.Println("host:", parsedURL.Host)
fmt.Println("path:",parsedURL.Path)
fmt.Println("Rawquesry:",parsedURL.RawQuery)


//modifying url components
parsedURL.Path ="/newPath"
parsedURL.RawQuery ="cookie_sync_completed=false"

// constructing a url string from url object
 newURL:= parsedURL.String()
 fmt.Println("'new url",newURL)




}
