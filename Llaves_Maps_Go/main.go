package main

import "fmt"

//los maps son herramientas que nos permiten asiganarle un valor a un elemento y este valor siempre se encontrara el mismo elemento.
func main() {
	//make no solo sirve para crear herramientas como maps, tambien nos permite crear diversas herramientas que se pueden encontrar en go doc.
	//dentro de los corchetes pondremos el tipo de valor que va a ir  allí para poder acceder al mismo mas adelante
	var m  = make(map[string]int)

	m["Jose"] = 20
	m["Juan"] = 21

	fmt.Println(m)
	

	//Recorrer un Map
	
	for i,v := range(m){
//De esta manera recorreremos tanto los elementos como los valores que poseen 
		fmt.Println(i,v)


	}

}
