//Map adalah salah satu tipe data, yang dimana datanya have a key yang unix. Key disini adalah index

package main

import "fmt"

func main() {
	var buah map[string]int //deklarasi
	buah = map[string]int{} //inisialisasi. Sesudah di deklarasi, jangan lupa di inisialisasi
	buah["Semangka"] = 7500
	buah["melon"] = 7500
	buah["buah naga"] = 13500
	fmt.Println(buah)
	fmt.Println("Harga buah semangka", buah["Semangka"])
	//kok indexnya string? karena diawal deklarasi tipe datanya string
	//Nama semangka dan melon adalah key-nya. Harganya adalah value.Jika key nya sama, maka yang diambil harga terakhir

	//delete()
	delete(buah, "melon")
	fmt.Println("Hasil setelah di delete buah melon:", buah)

	a := map[int]int{}
	a[0] = 1111 /*Value 1111 yang ada disini tergantung tipe data yang --> int{} sebelum kurung kurawal.
	Kalau tipe datanya integer, valueny harus integer*/
	a[1] = 2222 //key [0] dan [1] tergantung tipe data setelah map --> [int]
	//sama kek yang diatas (var buah map dst), tipe datanya string maka keyny string.
	//Valueny integer, maka nilainy harus integer. Tergantung deklarasinya seperti apa.

	fmt.Println(a)

	//cara deklarasi map yang lain adalah
	b := make(map[int]int)
	b[10] = 99
	b[11] = 100
	b[12] = 200

	delete(b, 12) //boleh dihapus karena berhubungan dengan value&isB
	fmt.Println(b)

	//map punya cara untuk mengecek key dalam map tsb. Misal ada 2 variabel seperti dibawah.

	value, isB := b[10]         //coba diganti indexny sesuai deklarasi map di atas
	fmt.Println("value", value) //value disini adalah mereturn nilai yang b[12]=200 yang diatas. Kan variabelny b
	fmt.Println("check value and it is", isB)
	/*kalau isB itu mereturn data booleans. Jika hasilnya true,
	maka valuenya ada. Jika false artinya ngga ada*/
}
