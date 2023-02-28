package main

import "fmt"

//funciones anonimas, estas se pueden modificar en tiempo de ejecucion
//declaramos una a variable con  el tipo funcion, le decimos que va a obtemer 2 parametros del tipo int
// y que va a devolver un tipo int
//para luego asignarle una funcion qie lleva la el tipado normal
//pero que no tiene nombre por que obtiene el nombre de la variable calculo
var Calculo func(int, int) int = func(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func main() {
	fmt.Println("suma de 2 numeros")
	fmt.Println(Calculo(2, 3))

	//al ser una funcion anonima y estar en una variable 
	//podemos redefinir la variable con otra funcion

	Calculo = func(num1 int , num2 int) int {
		return num1 - num2
	}

	fmt.Println(Calculo(2,3))
}


