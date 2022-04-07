package main

import "fmt"

//La funcion waitGroup es una mejor opcion cuando se requiere una eficiencia optima constante, por otra parte
//los channels son mucho mejores cuando no se requiere una eficiencia optima constante.

//Los channels tienen como funcion nativa la comunicacion entre las GoRoutines.

func say(text string,h chan string) {
	h <- text
}

func Duplicate(num int,f chan int){
	f <- num * 2
}
//Los channels permite comunicar la informacion de Un GoRoutine adicional con el GoRoutine main.
//de igual manera nos permite ejecutar mas de un GoRoutine que comparta un mismo channel.
func main() {

	c := make(chan string, 2)
	b := make(chan int,1)
	fmt.Println("Hello")

	go say("Bye",c)
	fmt.Println(<-c)

	go say("Hi",c)
	fmt.Println(<-c)

	go Duplicate(2,b)
	fmt.Println(<-b)

	
}
