package main
// Los slice son cambiables, mientras que los arrays no lo son.

import "fmt"

func main() {

	//Array
	//Se le debe indicar la cantidad de valores que deba de tener.
	var array [4]int
	array [0] = 1
	array [1] = 2
	fmt.Println(array, len(array), cap(array))
	// len nos permitara saber cuantos elementos hay en un array
	// cap nos permitira saber cuanta capacidad tenemos en un array

	//Slice
	//Nos permite no decirle la cantidad de valores que tendra.
	slice := []int{1,2,3,4,5,6} 
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el Slice
	// Los slices nos permitiran interactuar entre Slice, Array y Listas.
	
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	// el numero que va despues de los dospuntos es exclusivo
	// el numero que va antes de los dospuntos es inclusivo
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:]) 

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append nueva List
 	var newslice = []int{8,9,10}
	slice = append(slice, newslice... )
	//esto descomprimira la lista en la lista asignada.
	fmt.Println(slice)
}
