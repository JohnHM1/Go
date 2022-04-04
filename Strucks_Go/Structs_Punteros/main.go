package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

//de esta manera estaremos asociando el struct con una funcion
func (mypc pc) ping() {

	fmt.Println(mypc.brand, "pong")

}

//En esta funcion tomaremos el valor Ram de Struct pc y lo modificaremos
func (mypc *pc) duplicateRam() {
	mypc.ram = mypc.ram * 2
}

// Los punteros nos permiten el acceso a la memoria
//funcion main punteros.
func main() {

	a := 50
	//Este & nos permitira acceder a la direcion de memoria de la variable a y se la asignaremos a la variable b.
	b := &a

	fmt.Println(a)
	fmt.Println(&b)
	//Es * antes de la variable nos permitira acceder al valor que se guarda en esa direcion de memoria.
	fmt.Println(*b)
	//Como paradigma Estructurado siempre ira de lo superior a lo inferior
	//modificaremos el valor de a desde su direcion de memoria
	*b = 100
	fmt.Println(a)

	//Los punteros nos permiten trasladar Funciones/ Librerias de manera mucho mas simple.

	mypc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(mypc)

	mypc.ping()
	//impreciones de las modificaciones de los datos del struct atravez de funciones.
	fmt.Println(mypc)

	mypc.duplicateRam()
	fmt.Println(mypc)

	mypc.duplicateRam()
	fmt.Println(mypc)

}
