package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	address1 := Address{"Garut","Jawa Barat","Indonesia"}
	// pointer &
	address2 := &address1

	address2.City = "Bandung"

	// asterik
	*address2= Address{"Jakarta","DKI Jakarta","Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

}