package main

import "fmt"

func main() {
// condicional switch 
// se puede desarrollar de dos maneras distintas. tanto con la variable fuera del switch o dentro.
// se desarrolla cuando se va a iterar sobre una misma variable
	switch modulo := 4 % 2 ; modulo {

	case 0:
		fmt.Println("Es par")
	
	default:
		fmt.Println("Es impar")

	}
// switch sin condicion
// esta modolida se desarrolla cuando se itera con mas de una variable.
// o para anidar multiples condiciones tal cual lo hicieras con un if.
	var value int = -200
	switch {
	case value > 100:
		fmt.Println("Value es mayor que 100") 
	case value < 0:
			fmt.Println("Es menor a 0")

	default:
		fmt.Print("No condicion.")

	}
}
