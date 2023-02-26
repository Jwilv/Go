package main

import "fmt"

func main() {
	//creamos la variable de state y le aignamos un valor
	//luego le asignamos el state al switch
	switch state := 5; state {
	case 1:
		fmt.Println("es 1")
	case 2:
		fmt.Println("es 2")
	case 3:
		fmt.Println("es 3")
	case 4:
		fmt.Println("es 4")
	case 5:
		fmt.Println("es 5")
	default:
		fmt.Println("es mayor a 5")
	}
}
// evaluamos los casos 
