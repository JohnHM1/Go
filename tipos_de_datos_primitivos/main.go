package main 

import "fmt"

func main(){

	//Tener en cuenta que el lenguaje golang es de tipado estatico y dinamico combinado.
	//al declarar las variables de manera estatica ganaramos rendimiento 
	

	//Numeros enteros y su declaracion

	//int = Depende del OS (32 o 64 bits)
	//				puede contener valores desde / hasta.
	
	//      int8 = -128 a 127
	//   	int16 = -2**(15) a 2**(15)-1
	// 		int32 = -2**(31) a 2**(31)-1
	// 		int64 = -2**(63) a 2**(63)-1




	// Numeros enteros Positivos

	//uint = Depende del OS (32 o 64 bits)

	//      uint8 = 0 a 2**(8)-1
	//   	uint16 = 0 a 2**(16)-1
	// 		uint32 = 0 a 2**(32)-1
	// 		uint64 = 0 a 2**(64)-1


	
	// Numeros decimales 

	//float32 = 32 bits = +/- 1.18e**(-38) a +/- 3.4e**(38)
	//float64 = 64 bits = +/- 2.23e**(-308) a +/- 1.8e**(308)

	

	//Texto y Booleanos

	//string = ""
	//bool = true o false 



	//Números Complejos 


	//complex64 = Real e imaginario float32
	//complex128 = Real e imaginario float64

	// ejemplo : c := 10 + 8i

	var a complex128 = 10 + 8i

	fmt.Println("Esto es un numero imaginario: ",a)
}