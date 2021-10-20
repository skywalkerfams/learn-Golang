package main

import "fmt"

func main() {
	a := CetakNama("Setyo", "Aji")

	fmt.Println(a)
	fmt.Println(CetakNama("Setyo", "Aji"))
	b, c, d, e := Calculator(5, 5, 5) //b adalah eprtambahan, c adalah pengurangan. Sesuai urutan return dibawah
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func CetakNama(namaDepan, namaBelakang string) string { //ini adalah function yang mereturn value
	fullname := namaDepan + " " + namaBelakang
	return fullname
}

func Calculator(nilai1, nilai2, nilai3 int) (int, int, int, float64) {
	//Deklarasi tipe data jangan cuma 1,
	//karena return valuenya 2, tetapi deklarasinya cuma 1. Maka harus ditambah (int, int)
	pertambahan := nilai1 + nilai2 + nilai3
	pengurangan := nilai1 - nilai2 - nilai3
	perkalian := nilai1 * nilai2 * nilai3
	pembagian := float64(nilai1) / float64(nilai2) / float64(nilai3)

	return pengurangan, pertambahan, perkalian, pembagian //coba ditukar. Tergantung posisi return value.
}
