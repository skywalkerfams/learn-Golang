/* Array adalah kumpulan data yang disimpan dalam variabel yang dimana
array mempunyai kapasitas yang kita tentukan saat deklarasi.
Array akan dimulai indexnya dari 0 sedangkan lengthny dimulai dari 1.
Jangan sampai kebalik*/

package main

import "fmt"

func main() {
	var z int //jika tidak diberi value maka hasil runnya adalah 0 (default)
	fmt.Println(z)
	var y [7]int
	fmt.Println(y)
	/*[0 0 0 0 0 0 0] hasil runnya. Untuk mengganti value yang dimau, maka
	ingat, array memiliki index, maka buat seperti dibawah ini */
	var a [7]int

	a[0] = 1
	a[1] = 2
	a[3] = 3
	a[4] = 4
	a[5] = 5
	a[6] = 6
	fmt.Println(a) //Jika di run hasilnya akan seperti ini "[1 2 0 3 4 5 6]""

}
