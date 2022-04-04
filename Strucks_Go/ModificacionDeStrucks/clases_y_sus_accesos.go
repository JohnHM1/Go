package MyPackage

import "fmt"

//Existen dos tipo de clases las publicas y las privadas, Esto se hace con la Primera letra en Mayus para publica y min para privada.
type CarDos struct {
	Brand string
	Year  int
}

// PrintMessage imprimir mensaje
func PrintMessage(text string) {

	fmt.Println(text)
}
