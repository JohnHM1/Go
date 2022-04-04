package main

import (
	//Las letras a la derecha del directorio es un alias, con el cual se puede hacer llamado de esta Struck
	MyPackage "curso_golang_go/Strucks_Go/ModificacionDeStrucks"
	"fmt"
)

// Struck nos permitara crear una serie de variables dentro de una clase para luego ser asignadas a una
// variable.

//Estas tambien son denominadas estructuras de datos.
type car struct {
	brand string
	year  int
}

func main() {

	mycar := car{
		brand: "Ford", year: 1999,
	}
	fmt.Println(mycar)

	// La otra forma de hacer uso de un struck

	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2019

	fmt.Println(otherCar)
// Este es el llamado de Package de Un strucks Publico con la linea 5 de codigo.
	var carro MyPackage.CarDos
	carro.Brand = "Audi"
	carro.Year = 2022
	fmt.Println(carro)
}
