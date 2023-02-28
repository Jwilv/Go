package main

import "fmt"

//se inicializa con todas las posiciones con un 0, si no le asignamos nada
var tabla [10] int 

func main(){
	tabla[0] = 1
	tabla[5] = 6

	fmt.Println(tabla)
}