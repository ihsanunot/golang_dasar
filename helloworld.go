package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var a = 10
	var b = 13
	var c = a + b

	fmt.Println(c)

	var d = 10
	d += 10
	fmt.Println(d)

	var e = 9
	e++
	fmt.Println(e)

	var nama1 = "Ayana"
	var nama2 = "Shahab"

	namaLagi := "Mona Nabilah"
	namaLagi = "Neno Balqis"
	fmt.Println(namaLagi)

	var (
		namaDepan    = "Sarah"
		namaBelakang = "Azhari"
	)

	fmt.Println(namaDepan + namaBelakang)

	var namaLengkap bool = nama1 == nama2
	fmt.Println(namaLengkap)

	fmt.Println(len(nama1))
	fmt.Println("Shahab"[1])

	var konversiString = nama2[0]
	var hasilKonversiString = string(konversiString)
	fmt.Println(hasilKonversiString)

	var nilaiPertama int = 90
	var hasilNilaiUjian bool = nilaiPertama >= 90

	fmt.Println(hasilNilaiUjian)

}
