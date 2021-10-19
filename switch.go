package main

import "fmt"

func main() {
	a := 10 //coba diganti valueny antara 5 dan 10

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
}
