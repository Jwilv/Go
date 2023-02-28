package main

import "fmt"

func main() {
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

func uno(numero int) int {
	return numero * 2
}

//en este caso hacemos un return de 2 valores
// al ser mas de un tipo a retornar se pone ()

func dos(numero int) (int, bool) {
	if numero == 5 {
		return 10, true
	} else {
		return 0, false
	}
}

//en este caso al typo int de retornamos le asignamos un nombr e
//en este caso x
//en este caso poner un return seria suficiente para que retorne x
//que es como una variable de retorno

func tres(numero int) (x int) {
	x = numero * 3
	return
}

//ejemplo cuatro
//go al no tener sobre carga de funciones
//tenemos que indicarle que la cantidad de parametros que le vamos a enviar
//no esta totalmente definida
//ej:

//en este caso le decimos que todos los parametros que va a recibir son int
//pero que no sabemos cuantos
//entoces los almacena en numero

func cuatro(numero ...int) int {
	total := 0
	// en este caso obtenemos el rango total de numero
	//esto sirve para iterar uno por uno de los valores de numero
	//en primer parametro va a ser el index pero como no lo vamos a usar
	//ponemos un _
	//el segundo parametro es num, es donde se guarda el valor del numero actual
	for _, num := range numero {
		total = total + num
	}
	// el for se va a ejecutar durante todo el range de numero [ es como un map de js iteras por cada uno ]
	//y por cada numero se lo va a usar al total 
	//para luego de que termine le for retornar el total 
	return total
}
