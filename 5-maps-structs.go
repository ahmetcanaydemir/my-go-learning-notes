/*
Map: Key ve value tipi tutuyor
	sehirNufuslari := map[string]int{
		"Istanbul": 16000000,
		"Erzurum":  730000,
	}

	Map'te yeni bir key-value eklendiğinde sıra garanti edilmez

sehirNufuslari["Antayla"] = 555 -> ekleme güncelleme
delete(sehirNufuslari,"Antalya") -> mapten siler
pop, ok := sehirNufuslari["NotFoundCity"] -> ok değişkeni dizide olup olmadığını söyler
len(sehirNufuslari) -> map boyutunu verir


Struct:
type Doctor struct {
	number     int
	actorName  string
	comapnions []string
}

aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		comapnions: []string{
			"Liz shaw",
			"Jo Grant",
		},
	}

aDoctor.actorName

	aDoctor := struct{ name string }{name: "John Pertwee"} // anonim struct

	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"

	AnotherDoctor !== aDoctor olacaktır

	Inheritance:

	type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	veya
		b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	Aslında bu durumda Bird struct'ının Animal struct'ı ile hiçbir bağı yoktur. Sadece onun proplarını da
	yazmışız gibi davranır

	Tags
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

*/

package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
