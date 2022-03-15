//for es el unico ciclo iterativo.
package main

import "fmt"

func main(){

	//For condicional
	var i int
	for i = 0 ; i < 11 ; i++ {
		
		fmt.Println("esto es i:",i)

	}

	fmt.Printf("\n")

	//For While 
	var counter int 
	for counter < 11 {

		fmt.Println(counter)
		counter++

	}

	//For forever 
	/*var counterForever int
	for{
		fmt.Println(counterForever)
		counterForever++
	}
	*/

	var numero int = 10

	for numero > 0 {

		fmt.Println("Numero :", numero)
		numero--

	}
}