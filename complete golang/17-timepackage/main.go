package main

import (
	"fmt"
	"time"
)

//note - 02-01-2006,monday,15:04:05 yahi likhna hai har baar kyuki issi time pe golang  bani thi isliye unhone apna alag format bana diya jo ki yahi hai 
//note- 2006 represent year
//01- represent month
//02-reperesents day
// 15- represent hour
// 04-represent minute part
// 050 represent second part
//monday represents day of the week

func main(){
//jaise baaki lang mai hota haina dd-mm-yy or yyy-mm-dd golang mai aisa nhi hota iska alag hi format hai 
	// u have to write  this 02-01-2006 to get dd-mm-yyy

	//to get current time 
	currentTime:= time.Now()
	fmt.Println("time is :",currentTime)
	//but ye jo aayega na bada ganda format  mai hoga jisko hum read easily nhi kar skta
	// formatted:=currentTime.Format("dd-mm-yyy,day,hh:mm:ss")// aisa likhna se error aayega 
	formatted:=currentTime.Format("02-01-2006,monday,15:04:05") //dd-mm-yyy
	fmt.Println("formatted time ",formatted)


	//converting string to formatted time 
	datestr:= "2023-11-25"
	layoutstr:= "2006-01-02"
	formatted_time,_ := time.Parse(layoutstr,datestr)
	fmt.Println("formatted time ",formatted_time)

	//note- 2023-11-25 aisa hai isliye  2006-01-02 likha 
	//agar aisa hota 25/11/2030- 02/01/2006 aia likhta 

	//if u want to add one more day in current time 
	new_date := currentTime.Add(24*time.Hour)
	fmt.Println("new date:",new_date)
	formatted_new_date:=new_date.Format("2006-01-02")
	fmt.Println("formatted_new_date:",formatted_new_date)



}