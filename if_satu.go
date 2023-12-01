package main

import "fmt"

func main(){
	nama := "Ayana"

	if nama == "Ayana" {
		fmt.Println("Halo Ayana!")
	} else if nama == "Mona" {
		fmt.Println("Halo Mona!")
	} else {
		fmt.Println("Halo Guest")
	}
}