package main

import (
	pcmaster "curso_golang_go/Strucks_Go/reto"
	"fmt"
)

func main() {

	var dreamPc pcmaster.SuperPc
	dreamPc.MemoriaRam = 1000
	dreamPc.TarjetaGrafica = "3080ti"
	dreamPc.Board = "Intel"
	dreamPc.DiscoSSD = 2000

	fmt.Println(dreamPc)
}
