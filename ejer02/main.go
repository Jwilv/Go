package main

import "fmt"

//declaracion de variables
var numero int
var cadenaDeTexto string
var boolean bool

//al declarar el tipo pero no el  valor inicia con el valor por defecto 
//en caso de int es 0
//en caso de string es una cadena vacia ""
//en caso de bool es false 
//tambien le podemos asignar un valor 
// ej var numeroTest int = 10

//manera pro
// numero2 := 1

func main() {
	numero3, numero4, numero5, texto1 := 1, 2, 3, "cadena de texto"

	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)
	fmt.Println(texto1)
}
