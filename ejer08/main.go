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

	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Println(Calculo(2, 3))

	operacion()

	fmt.Println("ejemplo closures")

	//lo que pasa aca es que 
	//a miTabla no se le asigna Tabla, si no que 
	//se le asigna el valor que esta retorna
	//en este caso es una funcion, entonces Tabla se ejecuta una vez 
	//y esta guarda el estado 
	//por eso secuencia se graba 
	//por que solo se vuelve a ejecutar lo que retorna 

	miNumero := 5
	miTabla := Tabla(miNumero)

	for i:=1; i<=10; i++{
		fmt.Println(miTabla())
	}
}

//segundo ejemplo
func operacion() {
	resultado := func() int {
		a, b := 100, 30
		return a / b
	}

	fmt.Println("ejemplo 2 ")
	fmt.Println(resultado())
}

//ejemplo closures
//creamos una func que no devuelve una func que nos devuelve un int 
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia+=1
		return numero*secuencia
	}
}
