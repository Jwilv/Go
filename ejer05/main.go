package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int

var data string

func main() {
	fmt.Println("ingrese el numero1 : ")
	fmt.Scanln(&numero1)

	fmt.Println("ingrese el numero2 : ")
	fmt.Scanln(&numero2)

	fmt.Println(numero1)
	fmt.Println(numero2)

	//creamos el objeto scanner, con el estandar input que es el teclado

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ingrese el nombre de la operacion")

	//si se escanea algo se guarda en data 
	if scanner.Scan(){
		data = scanner.Text()
	}

	resultado := numero1 + numero2

	fmt.Println(data, resultado)

}

// "os"
// "bufio"
