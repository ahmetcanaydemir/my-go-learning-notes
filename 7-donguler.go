/*
For:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	Aynı anda i ve j atamak
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	s := []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// For range (foreach gibi)
	s := []int{1, 2, 3}
		for k, v := range s {
			fmt.Println(k, v)
		}
*/

package main

import (
	"fmt"
)

func main() {
	sehirNufuslari := map[string]int{
		"Istanbul": 16000000,
		"Erzurum":  730000,
	}
	for k := range sehirNufuslari {
		fmt.Println(k)
	}
}
