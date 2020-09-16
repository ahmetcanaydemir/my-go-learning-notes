/*

** Bool
- var n bool = true
- n := 1 == 1
- true veya false değerleri alabilir
- Default = false
- Integer alias'ı değil (0 = false değil)

** Integer
int type'i en minimum 32bit'tir ama sisteme göre otomatik genişler
var n int8 = 127 (signed)
var n uint16 = 42 (unsigned)

a & b = and
a | b = or
a ^ b = xor
a &^b = and not

a << 3 sola kaydır
a >> 3 sağa kaydır

** Floating Point
n := 3.14
	n = 13.7e72
	n = 2.1E14

** Karmaşık Sayılar
var n complex64 = 1+2i
var m complex128 = comlex(5,12)

gerçek kısım real(n)
sanal kısım imag(n)
	complex128 için float64, complex64 için float32

** String
s := "Bu bir string değişken"
s[2] -> 105 gibi bir sayı verir string(s[2]) harfi verir

s := "Bu bir string değişken"
	b := []byte(s) -> bize stringin byte'larından oluşmuş bir dizi verir

stringler + ile birleştirilebilir

** Rune
var r rune = 'a'
UTF32 kullanmak için (aslında int32'nin alias'ı)

*/

package main

import (
	"fmt"
)

func main() {
	var r rune = 'a'

	fmt.Printf("%v, %T\n", r, r)

}
