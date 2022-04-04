package main

import (
	"fmt"
	"strings"
)

func main() {

	var palabra string
	fmt.Scan(&palabra)
	isPalindromoUpperLower(palabra)

}

func isPalindromoUpperLower(text string) {
	//Con la funcion strings.To(Lower,Upper) se modifica el codigo ascii. (Mayus/Minus).
	text = strings.ToLower(text)
	var textoReversa string
	var i int
	for i = len(text) - 1; i >= 0; i-- {

		textoReversa += string(text[i])

	}
	if text == textoReversa {

		fmt.Println(true)

	} else {

		fmt.Println(false)

	}

}
