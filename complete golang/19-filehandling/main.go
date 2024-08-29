package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	"os"
)

func main() {
	//creating a file
	file, err := os.Create("example.txt")
	if err != nil {

		fmt.Println("error while creating ", err)
		return
	}
	defer file.Close() // this line is compulsory last mai chalegki sab kaam hona ke baad

	// //note- io.writestring is a conveninet way to write a string into file
	// 	content := "hello "// file jo create kari uske andar ka content
	// 	_,error:=io.WriteString(file, content ) // to insert content into file  t
	// 	if error!= nil {
	// 		fmt.Println("error while writing files :",error)
	// 		return
	// 	}

	// 	fmt.Println("succesfully created file ")

	//reading file - 2 ways - buffer and function
	//method 1
	//reading file using  buffer
	//data ko read karne ke liye hume pehle buffer create kaerna hota hai
	//buffer is like  a temporary storage

	// file,err:=os.Open("example.txt")
	// if err!= nil{
	// 	fmt.Println("error while opening  file",err)
	// 	return
	// }
	// defer file.Close()

	// //creating a buffer to read the file
	// buffer:=make([]byte,1024)

	// //read the file content into buffer

	// for {
	// 	int,err:=file.Read(buffer)
	// 	if err==io.EOF{ // read karte karte  end of file agya
	// 		break
	// 	}
	// 	if err!= nil{
	// 		fmt.Println("error while reading file ",err)
	// 		return
	// 	}

	// 	//process  the read content
	// 	//buffer mai jo bhi data hoga usko string mai convert kardenge
	// 	fmt.Println(string(buffer[:int]))
	// }

	//method 2 reading file using function ioutil

	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("errror while reading fiel ", err)
		return
	}
	fmt.Println(string(content))

}
