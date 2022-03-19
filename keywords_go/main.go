package main

import "fmt"

func main() {

	//defer, lo que realiza esta keyword , es , ejecutar primero las lineas de codigo que no tengan
	//defer como atributo a su derecha.
	// nos permite tener buenas practicas al momento del ingreso a base de datos , pues podemos
	// utilizarla para cerrar la conexion cuando finalice la funcion.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
	// una buena practica es usar un defer por funcion.

//Continue 
//Break
	
	for i := 0 ; i < 10 ; i++ {

		fmt.Println(i)
		//continue
		if i == 2 {

			fmt.Println("es 2")
			continue
			//continue nos permitira seguir ejecutando el proceso cuando se cumpla una condicion dada.
		}

		//break
		if i == 8 {

			fmt.Println("break")
			break
			//break nos permitira finalizar el procese cuando se cumpla una condicion dada.
		}

	}

}
