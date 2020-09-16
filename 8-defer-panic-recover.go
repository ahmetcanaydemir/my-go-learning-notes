/*
Defer:
Fonksiyon bitmeden çalışır timeout gibi

fmt.Println("start")
defer fmt.Println("middle")
fmt.Println("end")

Çıktı :
start
end
middle

defer fmt.Println("start")
defer fmt.Println("middle")
defer fmt.Println("end")

Çıktı :
end
middle
start

// Close yapmayı unutmamak için açıp hata kontrolü yapılıyor ve defer ile kapatılıyor
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()


	a:= "start"
	defer fmt.Println(a)
	a = "end"

	Çıktı:
	"start" -> çünkü o anki değerini saklıyor.


	Panic:
	Go'da exception bulunmuyor. Uygulama devam edemediği durumlarda panic oluşuyor

		a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

	Çıktı:
	panic: runtime error: integer divide by zero

	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Örneğin başka bir uygulama zaten bu porta bağlı
		panic(err.Error())
	}

	PANIC DEFER'den sonra gerçekleşir!! yani finally gibi kapanması gereken kaynağı kesinlikle kapatır

	fmt.Println("start")
	defer fmt.Println("defer")
	panic("panic")
	fmt.Println("end")

	Çıktı:
	start
	defer
	panic: panic

	Recover (bir nevi catch için)
	Eğer recover() nil değil ise uygulama panic durumundadır

	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("panic")
	fmt.Println("end")

	Çıktı:
	2020/09/14 21:54:13 Error: panic
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("panic")
	fmt.Println("end")
}
