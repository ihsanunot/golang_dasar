package main

import "fmt"

type Man struct {
	Nama string
}

func (man *Man) Married(){
	man.Nama = "Mr. " + man.Nama
}

func main(){
	ihsan := Man{"Ihsan"}
	ihsan.Married()

	fmt.Println(ihsan.Nama)
}