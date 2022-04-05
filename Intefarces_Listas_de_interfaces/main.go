package main

import "fmt"

//Las interfaces son un metodo que nos permite alojar muchos mas metodos dentro de sí.
//Este nos facilitara el acceso de datos a las funciones y a los valores de las structs.

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}
type rectangulo struct {
	base   float64
	altura float64
}

//"calcular": Esta funcion aloja la interfas que a su vez aloja la funcion area, misma que trabaja con dos structs diferentes.
//El argumento de ingreso sera destinado a lo alojado dentro de la interfas, que es la funcion area. Esta se ejecutara de manera distinta
//	en funcion al argumento de ingreso en la funcion calcular; alguna de las dos strucs.
func calcular(f figuras2D) {
	fmt.Println("El area:", f.area())

}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}
func main() {
	var micuadrado cuadrado
	var mirectangulo rectangulo

	micuadrado.base = 2
	mirectangulo.base = 2
	mirectangulo.altura = 3

	calcular(micuadrado)
	calcular(mirectangulo)

	// Lista de Interfaces

	//estas nos permitiran alojar en una lista varios datos que no pertenescan a un mismo tipo.

	var milista []interface{}

	milista = append(milista, "hola", 21, 2.0, true)
	fmt.Println(milista...)

}
