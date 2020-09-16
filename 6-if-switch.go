/*
	if true {
		fmt.Println("The test is true")
	}


	if pop, ok := sehirNufuslari["Erzurum"]; ok {
		fmt.Println(pop)
	}


	switch i := 2 + 2; i {
	case 1, 4:
		fmt.Println("bir veya 4")
	case 2:
		fmt.Println("iki")
	default:
		fmt.Println("not one or two")
	}


	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i değişkeni int")
	case float64:
		fmt.Println("i değişkeni float64")
	case string:
		fmt.Println("i değişkeni string")
	default:
		fmt.Println("i değişkeni başka bir tipte")
	}
*/

package main

import (
	"fmt"
)

func main() {
	var i interface{} = [3]int{}
	switch i.(type) {
	case int:
		fmt.Println("i değişkeni int")
	case float64:
		fmt.Println("i değişkeni float64")
	case string:
		fmt.Println("i değişkeni string")
	default:
		fmt.Println("i değişkeni başka bir tipte")
	}
}
