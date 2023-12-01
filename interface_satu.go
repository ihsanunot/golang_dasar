package main

import "fmt"

type HasName interface {
	// method-method
	GetName() string
}

func sayHello(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// implementasi Interface
func (person Person) GetName() string {
	return person.Name
}

func main(){
	person := Person{Name : "Anysa Syamsi"}
	sayHello(person)
}