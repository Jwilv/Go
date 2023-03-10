package main

import "fmt"

func exponencia(numero int) {
	if numero > 10000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}

func main() {
	exponencia(2)
}