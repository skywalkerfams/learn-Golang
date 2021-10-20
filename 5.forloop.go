/*Perulangan atau for loop adalah proses eksekusi untuk
mengulang suatu kode tanpa berhenti sampai kondisi yang kita tentukan tercapai*/

package main

import "fmt"

func main() {
	//for kemudian variabel nama bebas, lalu deklarasi bebas awalnya mau dari mana, contoh bawa skuy

	for a := 0; a < 5; a++ {
		fmt.Println("Nilai", a)
		/*	Nilai 0
			Nilai 1
			Nilai 2
			Nilai 3
			Nilai 4
			Nilai mulai dari 0-4, first we declared a mulai dari 0, yang dimana a akan dilooping sebanyak
			5x. Terus 0 lebih kecil dari 5 ngga? Yes, maka akan diprint nilai 0. Terus karena ada a++,
			a++ itu artinya maka a == 0 + 1.
			0 --> a = 0
			1 --> a = 0 + 1 = 1
			2 --> a = 1 + 1 = 2
			3 --> a = 2 + 1 = 3
			dst. Dikarenakan memenuhi syarat, maka hasil print selalu berulang sampai kondisi
			tertjapai. Terus kenapa cuma sampai 4? karena kan kurang dari 5 (lihat variabel) kecuali
			pakai lebih kecil sama dengan <=. Same like math :)*/

		//Dalam fungsi perulangan terdapat 2 statement.
		//Yaitu break dan continue. Contoh break
	}

	for b := 0; b <= 5; b++ {
		if b == 3 {
			break /* b lebih kecil dari 0, maka masuk ke baris ke 34 if b ==3. Jika b
			== 3, maka fungsi akan berhenti di 3 untuk menghentikan looping. Nah
			kan b nya 0, itu bukan == 3, maka bisa ke run dst sehingga b == 3. */
		}
		fmt.Println("Value dari b ", b)
	}
	//Ketika ada fungsi break, maka kode yang dibagian bawah akan berhenti, tidak ke run
	//makanya meskipun dilooping sampai 5, maka hanya akan dibreak sesuai dgn yang kita
	//tentukan, beda lagi dengan continue.

	for c := 0; c <= 5; c++ {
		if c == 3 {
			continue /* ketika c nya == 3, maka code will be stop at value 3.
			Setelah itu dilanjutkan lagi setelah memenuhi syarat*/
		}
		fmt.Println("Value dari c ", c)
	}

}
