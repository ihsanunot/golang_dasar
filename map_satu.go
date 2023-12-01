package main

import "fmt"

func main(){
	orang := map[string]string{
		"nama" : "Ayana Shahab",
		"alamat": "Tangerang Selatan",
	}

	fmt.Println(orang)
	fmt.Println(orang["nama"])
}