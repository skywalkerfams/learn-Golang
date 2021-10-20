package main

import "fmt"

func main() {
	fmt.Println("Ini tes ke Tiga save dari VSCode")
	fmt.Println("I want to be success")

	//Just a comment. Free to put anywhere and won't be show when run
	//fmt func can't used outside the func main()

	//func to print yes
	fmt.Println("Yes")

	//    /* words */ is a multiple comment.
	/*
		//func to print yes
		fmt.Println("Yes")
	*/

	/* Variabel adalah suatu perintah untuk menampung sebuah value
	contoh: var(spasi)nama_variabel(spasi)tipe_data(spasi)="x"
	*/

	var nama string = "Setyo Aji"
	var umur int = 22
	var bulin bool = true

	fmt.Println("My name is", nama, "and I'm", umur, "years old")
	fmt.Println("The statement above is", bulin)

	// Constanta atau line code : const harus memiliki nama_variabel yang berbeda
	const tempat_tinggal string = "Nuri Street No 15"
	fmt.Println("I lived in", tempat_tinggal)

	umur = 17
	fmt.Println("My name is", nama, "and I'm", umur, "years old")
	/*kenapa yang di print umur 17 dan bukan 22?
	karena sudah beda variabel (re-assign).
	Kasih beda contoh lagi*/

	/*Coba ini dihapus commentny terus run codeny.

	tempat_tinggal = "Nuri Street No 51"
	fmt.Println("I lived in", tempat_tinggal)


	/*Jalankan kode di atas maka akan muncul error. Cannot assign to tempat_tinggal (declared const)
	Bedanya const dan variabel, kalau const kalau sudah dideklarasi dan sdh diassign
	valuenya/isinya tidak boleh lagi dikasih value/assign yang berbeda. Kalau variabel
	bisa diganti/reassign valuen= atau nilainya. Itu aja bedanya. Kenapa umur bisa terganti?
	karena umur itu variabel awal deklarasinya. */

	//coba dihapus tempat_tinggal yang nuri street no 51. Lalu run
}
