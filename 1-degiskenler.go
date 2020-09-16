/*

Değişken oluşturmanın 3 yolu

- var i int = 42

- var j int
  j = 33

- z := 44

Go'da tanımlanan tüm değişkenler kullanılmak zorunda

Değişkenler yeniden tanımlanamaz ama scope içinde shadow yapılabilir

Değişken görünürlüğü
- Büyük harfle başlayan global değişkenler Export edilir
- Küçük harfle başlayan global değişkenler package da scoplanır (namespace gibi)
- private yok

Tip dönüşümü:
- float32(int_sayi) örneğindeki gibi type fonksiyonmuş gibi kullanılır
- string'e dönüştürmek için strconv paketi gerekir
*/

package main

import (
	"fmt"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = i
	fmt.Printf("%v, %T\n", j, j)

}
