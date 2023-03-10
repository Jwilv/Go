package main

import "fmt"

func main() {

}

func exponencia(numero int) {
	if numero < 10000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}