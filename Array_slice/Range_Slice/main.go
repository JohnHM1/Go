package main

import (
	"fmt"
)

func main() {

	var slice = []string{"hola", "que", "hace"}

	//La funcion range nos permitira ver dos conceptos de nuestros slice, Listas y Arrays; La posicion y el contenido de esa posicion.

	var elemento int
	var contenido string

	for elemento, contenido = range slice {

		fmt.Println(elemento, contenido)

	}
	//Ejecutando la Funcion.
	isPalindromo("ama")
}

//
// Actividad ; Detectar si la palabra es palindroma o no.
func isPalindromo(text string) {

	var textReverse string
	// La i tomara el valor del numero de caracteres que tenga la palabra.
	// Se le resta 1 con el proposito de contabilizar la casilla numero 0.
	for i := len(text) - 1; i >= 0; i-- {
		//se le asigna a la variable el caracter que se encuentre en la casilla numero i. Se transforma de esta manera el codigo ascii.
		textReverse += string(text[i])
	}
	//si la palabra ingresada es igual a la variable : true, si no , false.

	if text == textReverse {

		fmt.Println(true)

	} else {

		fmt.Println(false)

	}

}
