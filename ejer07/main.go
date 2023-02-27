package main

import "fmt"

func main(){
fmt.Println(uno(5))

//vamos a crear 2 varaibles por que la funcion devuelve 2  valores 
numero, estado := dos(1)
println(numero)
println(estado)
}

//estructura de una funcion 
// declaramos la funcion con func 
//luego el nombre de la funcion
//luego los parametros si es que tiene y su tipo
// luego el tipo a retornar puede no retornar nada

func uno( numero int) int {
return numero*2
}

//en este caso hacemos un return de 2 valores 
// al ser mas de un tipo a retornar se pone ()

func dos(numero int) (int, bool){
	if numero == 5{
		return 10, true 
	}else{
		return 0, false 
	}
}