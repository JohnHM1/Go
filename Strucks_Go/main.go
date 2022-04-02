package main

import (
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

	// La otra forma de hacer un struck

	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2019

	fmt.Println(otherCar)

	var carro MyPackage.CarDos
	carro.Brand = "Audi"
	carro.Year = 2022
	fmt.Println(carro)
}
