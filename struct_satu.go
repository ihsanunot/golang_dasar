package main

import "fmt"

// struct 
type Customer struct {
	Name, Address string
	Age int
}

// struct method
func (customer Customer) sayHello(){
	fmt.Println("Halo saya lagi coba Struct Method dari", customer.Name)
}


func main(){
	var namaCustomer Customer
	namaCustomer.Name = "Ayana Shahab"
	namaCustomer.Address = "Jakarta"
	namaCustomer.Age = 26

	fmt.Println(namaCustomer)

	// struct kedua
	contohStruct := Customer{
		Name : "Mona Nabilah",
		Address : "Indonesia",
		Age : 28,
	}

	fmt.Println(contohStruct)

	// struct method
	neno := Customer{Name: "Neno"}
	neno.sayHello()
	
}