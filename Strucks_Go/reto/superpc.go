package main

import "fmt"

type SuperPc struct {
	TarjetaGrafica string
	MemoriaRam     int
	Board          string
	DiscoSSD       int
}

func (miPc *SuperPc) MasRam() {

	miPc.MemoriaRam = miPc.MemoriaRam * 2

}

func main () {
	var miPc SuperPc

	miPc.Board = "Asus"
	miPc.DiscoSSD = 1000
	miPc.TarjetaGrafica = "3080ti"
	miPc.MemoriaRam = 16
	//Quiero mas Ram
	//Mi memoria Ram antes:
	fmt.Printf("Mi memoria ram es: %d\n", miPc.MemoriaRam)
	miPc.MasRam()
	//Mi memoria Ram despues:
	fmt.Printf("Mi memoria ram es: %d\n", miPc.MemoriaRam)

}
