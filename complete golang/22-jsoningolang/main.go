//inm go ,the encoding/json package is used to encode and decode json(javascript object notation)
//json is a lightweight data interchange format
//struct ,maps,slicves  ,or any type of datas ko  ko achaa format mai daalna ka tareeka json hai taaki isko hum share kar payaa
//JSON is not just for structs but can be used for almost any type of data in Go that you might need to share or store in a structured format. It’s a flexible and widely-used format, especially for web applications where data needs to be sent back and forth between a server and a client (like a web browser).
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{Name: "John", Age: 30}
    fmt.Println("person data is:", person)

    // Convert person into JSON encoding (marshalling)
    jsondata, err := json.Marshal(person)
    if err != nil {
        fmt.Println("error marshalling to JSON:", err)
        return
    }

    fmt.Println("JSON data is:", string(jsondata))

    // Decoding (unmarshalling)
    var personData Person
    err = json.Unmarshal(jsondata, &personData)
    if err != nil {
        fmt.Println("error unmarshalling:", err)
        return
    }

    fmt.Println("person data is:", personData)
}
