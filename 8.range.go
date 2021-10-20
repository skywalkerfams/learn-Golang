//range kurang lebih sama dengan loop, yaitu pengulangan.
//Biasanya ini dipakai untuk data bertipe array, slice atau map.
//Terkadang saat ngoding we just need the elements, not the	Index.
//Bedanya range ama forloop, range menggunakan balikan 2 data. Index dan elemen.
package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d", "e"}

	fmt.Println(a)
	fmt.Println("<==================================>") //Pembatas
	//kalau pakai loop bakal kek gini
	//for i := 0; i < count; i++ {} //yang dimana countnya ini diganti dengan panjang dari sebuah array si len(a)
	for i := 0; i < len(a); i++ { //yang dimana dimulai dari value 0,
		fmt.Println(a[i]) //maka a index ke i. a index ke 0, a index ke 1 adalah b dst.
	}
	fmt.Println("<==================================>") //Pembatas
	//Yang tadi menggunakan loop, terus kalau range?

	for _, value := range a { //awalnya index, lalu diganti jadi underscore, coba ganti lagi _ menjadi index
		//index disini sama seperti for loop yang i := 0.
		//Value ini adlaah sebuah elemen yang dimana kalau kita sebelumnya memakai for-loop, itu harus dengan cara
		//a[i] atau a index ke i itu hasilnya berapa. Jika pakai range, kita tidak perlu pakai a index ke berapa, tapi
		//dengan begini, kita ambil valuenya aja, program mengerti valuenya a, b, c dst.
		//fmt.Println("Index menggunakan range", index) //coba dikomen line ini lalu akan muncul garismerah di index, arahkan cursor kstu
		fmt.Println("Value menggunakan range", value)
	}

	fmt.Println("<==================================>") //Pembatas
	//atau jika ingin memakai range dengan cara pendek maka hasilnya seperti ini.

	for index := range a {
		fmt.Println("Value menggunakan range2:", a[index])
	}

	/*Dirange, bisa direturn 2 value atau mengembalikan 2 data. for index, value := range a {}
	data yang pertama berupa index nilainya, 0,1,2 dst. Sedangkan data ke dua yaitu value
	nilai balikannya adalah sebuah value dari array/slice/map misal a atau b atau c dst. Tergantung
	valuenya apa. Bisa string, boolean dll. */
	//Kalau mau pakai data tapi indexnya gk dibutuhkan, ganti dengan undescore _ seperti dibawah
	//for _, value := range a {}

	fmt.Println("<==================================>") //Pembatas
	fmt.Println("<contoh lain>")                        //Pembatas

	b := map[string]int{"mobil": 10, "motor": 5, "hp": 2}
	fmt.Println(b)

	for i, v := range b {
		fmt.Println("Index:", i)
		fmt.Println("value:", v)
	}
	fmt.Println("<=====>") //Pembatas
	c := map[int]int{0: 10, 1: 5, 2: 2}
	fmt.Println(c)

	for _, v := range c { //index i diganti underscore jika tidak ingin dipakai
		//fmt.Println("Index:", i) //hilangkan komen, ganti _ jadi i untuk mengetes kode
		fmt.Println("value:", v)
	}
}
