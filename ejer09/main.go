package main

import "fmt"

//se inicializa con todas las posiciones con un 0, si no le asignamos nada
// ponemos el nombre, luego las posiciones, y luego el tipo que va a almacenar 
var tabla [10] int 

func main(){
	tabla[0] = 1
	tabla[5] = 6

	fmt.Println(tabla)

	//segunda forma de iniciarlo 
	fmt.Println("ejemplo 2")
	tabla2 := [10] int{1,2,3,4,5,6,7,8,9,10}
	//len no da la longitud del array 
	for i:=0; i<len(tabla2); i++{
		fmt.Println(tabla2[i])
	}

	//ejemplo de matrizes 
	var matrizes  [5][7]int
	matrizes [2][1] = 57
	fmt.Println(matrizes)

}