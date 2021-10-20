//function itu adalah kumpulan statement yang dijalan bersamaan dalam 1 task
//function dibagi 2
//function yang bisa mengembalikan value, dan yang tidak sama sekali

package main

import "fmt"

func main() {
	var a string = "Aji"
	fmt.Println(a)
	//semua coding berada dalam func main.
	//kemudian dibawah adalah contoh function yang tidak mereturn sama sekali

	printname()
	fmt.Println(printnamereturnvalue()) //jadi ini mencetak diluar func main
	printnameusingparameter("SETYO")
	fmt.Println(printnameusingparameterreturnvalue("IMAM"))
	Pertambahan(10, 5)
}

func printname() {
	var nama = "Saim"
	fmt.Println(nama)
} //terus kalau dirun maka gk muncul nama saim, karena ketika sudah membuat func, func harus dipanggil
//atau dimasukkan ke dalam func main () {}. Caranya adalah mencopy printname() lalu paste ke dalam
//func main seperti di atas. baru saim bisa muncul karena ditaruh di dalam func main.

func printnamereturnvalue() string {
	var nama = "SAIM"

	return nama
}

//ini function yang masih polos. Belum ditambah parameter.
func printnameusingparameter(nama string) {
	fmt.Println(nama)

}

//Ini beda function yah, jadi nama return itu dari (nama string) string {}
func printnameusingparameterreturnvalue(nama string) string {
	return nama

}

//Digolang itu kalau nama atau huruf depannya besar maka itu public
//kalau huruf depannya kecil itu private. Coba dirun

func cetaknamaprivate() {
	var nama = "Saim"
	fmt.Println(nama)
}

func Pertambahan(a, b int) { //bisa juga (a, b int) kalau tipe datanya sama
	fmt.Println("Hasil pertambahan a/b =", a/b) //+, -, *, /, %
}
