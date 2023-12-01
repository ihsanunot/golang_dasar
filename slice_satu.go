package main

import "fmt"

func main(){
	names := [...]string{"Ayana","Mona","Neno","Sarah","Ann","Litch"}
	slice := names[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
}