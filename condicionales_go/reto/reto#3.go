package main

import (
	"fmt"
	"log"
	"strconv"
)


func main() {
	
	var numero int = 2
	numero_Par_Impar(numero)
	var nombre string = "John"
	var clave string = "1234"
	comparador_usuario_contraseña(nombre, clave)


}




func numero_Par_Impar(numero int) {

	if numero%2 == 0 {

		fmt.Println(true)

	} else {

		fmt.Println(false)
	}

}

func comparador_usuario_contraseña(nombre_usuario, contraseña_usuario string) {

	if nombre_usuario == "John" {
		fmt.Println(true)
		contraseña_usuario, err := strconv.Atoi(contraseña_usuario)
		if err != nil {
			log.Fatal(err)

		} else {

			if contraseña_usuario == 1234 {

				fmt.Println(true)
				fmt.Println("acceso Permitido.")
			}
		}
	}
}


