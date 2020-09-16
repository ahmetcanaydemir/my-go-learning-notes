/*

Açılış süslü parentezi aynı satırda, kapanış olan tek olmak zorunda
func main() {
	fmt.Println("test")
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

// iki parametre de string olduğundan string yazmaya gerek yok
func sayMessage(greeting, name string) {
	fmt.Println(greeting, name)
}

// Pointer parametre
func main() {

	greeting := "Hello"
	name := "Stacey"
	sayMessage(&greeting, &name)
	fmt.Println(name)

}

func sayMessage(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// ... arguments
func main() {

	sum(1, 2, 3, 4, 5)

}

func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

// Pointer ve return

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *s)
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}


//Return değişkenini- isimlendirme
func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// Error ile birlikte return etmek
func main() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Sıfıra bölünemez")
	}
	return a / b, nil
}


// Anonim fonksiyon
	func () {
		fmt.Println("Hello go!")
	}()


		for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}


		f := func() {
		fmt.Println("Hello go!")
	}
	f()



	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("0'a bölme hatası")
		}
		return a / b, nil
	}
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// Methods (Structlar için extension veya class methodu gibi)
	// Tek farkı fonksiyonun sol tarafında parantez içerisinde struct nesnesi var
	func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}


*/

package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet("Selam")
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet(newName string) int {
	fmt.Println(g.greeting, g.name, newName)
	g.name = "Selam"
	return 5
}
