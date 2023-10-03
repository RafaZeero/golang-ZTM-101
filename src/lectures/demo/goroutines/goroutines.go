package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}

	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("Done: %c \n", r)
	}

	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}

	time.Sleep(2 * time.Second)

	fmt.Printf("Capitalized: %c \n", capitalized)
}

// // My implementation
// func main() {
// res := []rune{}
// go func() {
// 	for _, r := range data {
// 		res = append(res, Capitalize(r))
// 		fmt.Printf("Current strings: %v\n", res)
// 	}
// }()

// time.Sleep(1 * time.Second)

// fmt.Println(res)
// }

// func Capitalize(l rune) rune {
// return unicode.ToUpper(l)
// }
