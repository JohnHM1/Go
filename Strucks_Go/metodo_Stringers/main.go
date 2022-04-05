package main

import "fmt"

type pc struct {
	tarjetagrafica string
	ram            int
	disk           int
}
//Esta funcion trabajara sobre la variable y la struct definidas en su principio.
//Luego se escribira la funcion "String" y despues el tipo de dato que retornara.
func (miPc pc) String() string {
	return fmt.Sprintf("Tengo %dGB de RAM, %dGB de disco y una targeta grafica %s", miPc.ram, miPc.disk, miPc.tarjetagrafica)

}

//con el metodo stringers podremos modificar los output's de las structs en consola.

func main() {
	miPc := pc{ram: 16, tarjetagrafica: "3080ti", disk: 1000}
	fmt.Println(miPc)
//De esta manera personalizaremos nuestro output en consola.
}
