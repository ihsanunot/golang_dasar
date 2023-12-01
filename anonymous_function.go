package main

import "fmt"

type BlackList func(string) bool

func registerUser(nama string, blackList BlackList){
	if blackList(nama) {
		fmt.Println("anda di block", nama)
	} else {
		fmt.Println("Welcome", nama)
	}
}

func main(){
	blackList := func (nama string) bool {
		return nama == "Mona"
	}

	registerUser("Ayana", blackList)
}