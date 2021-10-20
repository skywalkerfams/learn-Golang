/* Array adalah kumpulan data yang disimpan dalam variabel yang dimana
array mempunyai kapasitas yang kita tentukan saat deklarasi.
Array akan dimulai indexnya dari 0 sedangkan lengthny dimulai dari 1.
Jangan sampai kebalik.*/

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

	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	a[5] = 60
	a[6] = 70
	fmt.Println(a) //Jika di run hasilnya akan seperti ini "[10 20 30 40 50 60 70]"

	var counterpain [5]string //nama variabel bisa bebas, kebetulan lg pakai counterpain

	counterpain[0] = "Methyl salicylate"
	counterpain[1] = "Eugenol"
	counterpain[2] = "Menthol"
	counterpain[3] = "Vanishing cream base to"
	counterpain[4] = "Analgesic cream"
	fmt.Println(counterpain)

	fmt.Println(counterpain[1])
	fmt.Println(counterpain[2])
	fmt.Println(counterpain[4])
}
