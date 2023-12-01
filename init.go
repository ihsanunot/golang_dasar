package main

import (
	"fmt"
	"belajar-golang-dasar/database"
	//contoh blank indefitier
	_ "belajar-golang-dasar/internal"
)

func main(){
	// package + function
	fmt.Println(database.GetDatabase())
}