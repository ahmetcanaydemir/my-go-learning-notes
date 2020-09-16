/*
	grades := [3]int{97, 85, 93}
	grades := [...]int{97, 85, 93} (verilen eleman kadar boyut oluştur)
	var students [3]string

	len(students) boyutu verir

	2 Boyutlu dizi:
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}

	Slices:
	Slices kullanırken pointer olmadan atama yapılabilir atama yapılan obje aynı objedir
	a := []int{1, 2, 3}
	len(a)
	cap(a)


	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // Tüm elemanları al
	c := a[3:]  // sondan 4 elemanı al
	d := a[:6]  // ilk 6 elemanı al
	e := a[3:6] // 4., 5. ve 6. elemanı al

	make([]int, 3) -> 3 elemanlı slice oluştur [0 0 0]
	append(a, 1) -> eleman ekleme

	[]int{2, 3, 4, 5, 6}... ->sondaki 3 nokta javascript'teki gibi çalışır
*/

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a)

}
