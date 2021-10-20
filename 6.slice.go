/* Slice adalah referensi elemen dari array.
Kalau slice bisa diresize, sedangkan array tidak */
package main

import "fmt"

func main() { //contoh array
	var array [3]int
	array[0] = 23

	array1 := [3]int{25, 26, 27}
	fmt.Println(array)
	fmt.Println(array1)

	slice := []int{99, 100, 101, 102} //Slice ketika pertama kali dideklarasi, tidak perlu length. Ini cara pertama slice
	slice2 := make([]int, 2)          //coba diganti2 nilai lengthny. Make punya 2 parameter type ama size
	fmt.Println(slice)
	fmt.Println(slice2)

	//ini cuma spasi

	fmt.Println("Panjang Array", len(array)) //fungsi len adalah untuk mengetahui panjangnya dari sebuah array
	fmt.Println("Panjang Array1", len(array1))
	fmt.Println("Panjang Slice", len(slice)) //len juga akan mengikuti panjangnya array meskipun awalnya gk diinput

	slice = append(slice, 1000)
	slice = append(slice, 1001)
	slice = append(slice, 1002) //append berfungsi menambah index
	slice = append(slice, 1003)
	slice = append(slice, 1004)
	slice = append(slice, 1005)
	fmt.Println(slice)

	slice2[0] = 1               //[1 0 33 34 35 36] karena index 1 gk ada valueny maka 0
	slice2 = append(slice2, 33) //2
	slice2 = append(slice2, 34) //3
	slice2 = append(slice2, 35) //4
	slice2 = append(slice2, 36) //5
	fmt.Println(slice2)

	fmt.Println(slice2[3])

	//hapus komen ini dan sisakan array sama dengan lalu run untuk tes -> array = append(array, 99)
	//ada error jika cursor diarahkan ke array. Invalid argumen. Array gk bisa diappend.

	// for i := 0; i < count; i++ {
	//i akan dilooping sebanyak count. Count itu bentuknya integer, gk bisa float (koma)
	//mod atau % adalah nilai sisa bagi. 3%2=1, 6%2=0, dst.

	/* for i := 0; i < len(slice2); i++ {
		fmt.Println("Looping results", slice2[i])
	} */ //multiple comment ini bisa dihapus

	for i := 0; i < len(slice2); i++ {
		if slice2[i]%2 == 0 {
			fmt.Println("Hasil Perulangan Genap", slice2[i])
		} else {
			fmt.Println("Hasil Perulangan Ganjil", slice2[i])
		}
	}

}
