package main

import "fmt"

var state bool 

func main(){

	state = true 

	//los parentecis () no son necesario al menos que
	//se necesite por ej
	//hacer un calculo matematico 
	// (1+2)*(2+3)

	if state == true {
		fmt.Println(state);
	}else{
		fmt.Println("es falso")
	}

	//se sigue la misma regla, el inicio del corchete tiene que estar en la misma linea

	//se le puede asignar un valor a una variable en la misma linea del if
	// y luego condicionarlo 
	if state=false; state == false{
		fmt.Println("se asigno")
	}

}