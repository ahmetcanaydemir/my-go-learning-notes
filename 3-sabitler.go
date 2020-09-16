/*
Sabit değişeknler
	const myConst int = 4

const myConst float64 = math.Sin(1.57)
gibi runtime sırasında hesaplanması gereken değerler constant olamaz

değiştirilemezler ama shadow olabilir

tip vermeden const a = 42 kullanılabilir

	const a = 42
	var b int16 = 27
	a+b
	bu durumda a int16 gibi davranacaktır

Aşağıdaki durumda compiler bizim için aynı şekilde devam ettiriyor
const (
	a = iota
	b
	c
)
iota enumerator gibidir scope içinde 0 1 2 gibi değerler alır
enum oluşturmak için mantıklı

const (
	_ = iota
	kedi
	kopek
	yilan
)
	var hayvanTipi int = kedi
	fmt.Printf("%v\n", hayvanTipi == kedi)

_ işareti c#'taki gibi değer önemsiz anlamına gelir



*/

package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isTechnician

	canSeeEurope
	canSeeAfrica
	canSeeAsia
)

func main() {
	var roles byte = isAdmin | canSeeEurope | canSeeAsia

	fmt.Printf("%b\n", roles)
	fmt.Printf("Admin mi? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Teknisyen mi? %v", isTechnician&roles == isTechnician)

}
