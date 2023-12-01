package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func gantiNegara(address *Address){
	address.Country = "Amerika"
}

func main(){
	address := &Address{}
	gantiNegara(address)

	fmt.Println(address)
}