package main

import "fmt"

func main(){
	var names [3] string
	names[0] = "Ayana"
	names[1] = "Shahab"
	names[2] = "Ihsan"

	fmt.Println(names[0])

	fmt.Println("==================================")

	var nilai = [...] int{
		90,
		80,
		70,
		60,
		50,
		40,
	}

	fmt.Println(nilai)
	fmt.Println(nilai[0])
	fmt.Println(len(nilai))
}