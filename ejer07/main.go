package main

import "fmt"

func main(){
fmt.Println(uno(5))
}

//estructura de una funcion 
// declaramos la funcion con func 
//luego el nombre de la funcion
//luego los parametros si es que tiene y su tipo
// luego el tipo a retornar puede no retornar nada

func uno( numero int) int {
return numero*2
}

