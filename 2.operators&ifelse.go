package main

import "fmt"

func main() {
	fmt.Println("Learn operators")

	/* if statement
	==
	!=
	>=
	<=
	>
	<
	*/

	nama := "Setyo"

	if nama == "Setyo" {
		fmt.Println("Yes his name is true")
	} else {
		fmt.Println("Not same with the variables")
	}
	/*Karena disini saya punya variabel yang dimana nama := "Setyo"
	== if nama Setyo maka hasilnya pasti benar. Jika salah, maka
	akan diproses di else. Sebagai latihan, bisa ganti value dari if nama == "Aji"
	lalu save kemudian run lagi codenya. Maka akan diproses di else
	*/
	/*Untuk menggunakan data types tipe string, valuenya harus sama.
	begitu juga integer tipe datanya harus integer, boleans dgn boleans,
	float dan float. Tipe data harus sama untuk membandingkan sesuatu.

	/* coba dihapus multiple comment yang ini
		if 1 == "1" {
			fmt.Println("Apakah sama?")
		}
		maka ketika di run hasilnya:
		invalid operation: 1 == "1" (mismatched types untyped int and untyped string)
	*/

	nilai1 := 20
	nilai2 := 30

	result := nilai1 + nilai2

	if result == 50 {
		fmt.Println("Same numeric")
	} else if result > 50 {
		fmt.Println("Bigger than 50")
	} else {
		fmt.Println("Less than 50")
	}

	/*Jika if == sama, proses stop. Jika gk memenuhi syarat
	maka yang dibaca else if, jika gk memenuhi syarat maka pindah
	lagi ke bawah dst. */
}
