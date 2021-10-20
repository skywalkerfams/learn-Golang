package main

import "fmt"

func main() {
	a := 3 //coba diganti valueny antara 5 dan 10

	if a == 10 {
		fmt.Println("Nilai a adalah 10")
	} else if a == 5 {
		fmt.Println("Nilai a adalah 5")
	} else {
		fmt.Println("Nilai a bukan 10 ataupun 5")
	}
	/* command switch sama dengan if else.
	switch expression {
	case condition:
	*/
	switch a {
	case 10:
		fmt.Println("Nilai a menggunakan switch ad/ 10")
	case 5:
		fmt.Println("Nilai a menggunakan switch ad/ 5")
	default:
		fmt.Println("Nilai a menggunakan switch bukan 10")
	}

	if a == 5 {
		fmt.Println("Nilai a = 5")
	} else if a > 5 {
		fmt.Println("Nilai a lebih besar dari 5")
	} else {
		fmt.Println("Nilai a kurang dari 5")
	}

	b := 2
	if b >= 0 || b%3 == 0 {
		fmt.Println("Nilai b habis dibagi 3")
	} else {
		fmt.Println("Nilai b tidak habis dibagi 3")
	}
	/* nilai b adalah 2 lebih besar dari 0, iya.
	2 habis dibagi 3? 2 gk habis dibagi 3, 2 habis dibagi 2.
	Kenapa saat dirun 2 larinya ke "Nilai b habis dibagi 3"?
	Karena ada atau nya || yang dimana jika disalah satu kondisi
	yang pertama b>= 0 atau b%3 itu true atau memenuhi syarat, maka
	hasilnya akan masuk di if statement 1 yang nilai b habis dibagi 3.
	Jika secara logika, nilai 2 tidak habis dibagi 3. Untuk itu kita
	harus hati2 dalam memakai simbol kondisi. Maka kita ganti dengan && dan dan */

	c := 2 //variabelnya diganti c karena a dan b sudah terpakai di atas
	if c >= 0 && c%3 == 0 {
		fmt.Println("Nilai c habis dibagi 3")
	} else {
		fmt.Println("Nilai c tidak habis dibagi 3")
	}
}
