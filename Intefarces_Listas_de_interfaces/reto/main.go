package main

import (
	pb "curso_golang_go/Intefarces_Listas_de_interfaces/reto/metodos_y_structs_Publicos"
	"fmt"
)

type areas interface {
	Area1() float64
}

func calcularPrivado(f areas) {
	fmt.Println("El area es:", f.Area1())
}
func main() {
	var cuadradoPu pb.Cuadrado
	var rectanguloPu pb.Rectangulo
	var circuloPu pb.Circulo

	cuadradoPu.Base = 2
	rectanguloPu.Altura = 3
	rectanguloPu.Base = 3
	circuloPu.Pi = 3.14
	circuloPu.Radio = 4

	calcularPrivado(cuadradoPu)
	calcularPrivado(circuloPu)
	calcularPrivado(rectanguloPu)
}
