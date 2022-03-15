package main

import (
	"fmt"
	
)

func main(){

	// Declaracion variables
	var hello  string = "Hello"
	var helloworld  string = "world"

	//Println

	fmt.Println(hello,helloworld)

	//Printf

	var nombre string = "Platzi"
	var cursos int = 500
/////////////TIPOS DE PARCEO
	//booleano: %t
	//int, int8 etc.: %d
	//uint, uint8 etc.: %d, %#x si se imprime con %#v
	//float32, complex64, etc: %g
	//cadena: %s
	//array: %p
	//slice: %p

	fmt.Printf("%s tinene más de %d cursos\n", nombre , cursos)

	//Si no conoces el tipo de dato que se deberia imprimir puedes usar %v que sera el valor predeterminado en la variable.

	fmt.Printf("%v tinene más de %v cursos\n", nombre , cursos)

	//Sprintf

	//este comando nos permitira transformar el output sin que salga en la terminal, en un string que puede ser guardado en una variable.

	var message string = fmt.Sprintf("%s tinene más de %d cursos", nombre , cursos)
	fmt.Println(message)


	//Conocer el tipo de dato 

	fmt.Printf("hello: %T\n",hello)
	fmt.Printf("cursos: %T\n",cursos)

	/* esto es un comentario */

	

	


}