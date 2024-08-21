package main

import "fmt"
//switch statement
func main() {
	day := 8
	switch day {
	case 1,30:
		fmt.Println("monday")
	case 2,22:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")

	default:
		fmt.Println("unknown day")
	}

//switch statement with multiple values

	month:="march"
	switch month{
	case  "january ", "february" ,"december":
		fmt.Println("winter")
	case "march", "june":
		fmt.Println("summer")
	default:
		fmt.Println("other season")

	}

}
