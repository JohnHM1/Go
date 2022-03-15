package condicional

import (
	"fmt"
	"log"
	"strconv"
)

func condicional() {

	var valor int = 1
	var valor_2 int = 2

	if valor == 1 {

		fmt.Printf("Es un %d\n", valor)
	} else {

		fmt.Println("No es 1")

	}

	//Con el operador logico AND

	if valor == 1 && valor_2 == 2 {

		fmt.Println("Es verdad")

	}

	//Con el operador logico OR

	if valor == 0 || valor_2 == 2 {

		fmt.Println("Es verdad, OR")

	}

	//Convertir texto a numero

	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value:", value)

	
}
