package main

import "fmt"

func main() {
	//distintas formas del for
	//ejemplo 1
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//segundo ejemplo 
	//forma tradicional 
	//no rompe con la otra i por que esta esta dentro del scope del for 

	for i:=0; i <= 10; i++ {
		fmt.Println("i es una variable local dentro del for")
	}

	//tercer ejemplo
	//bucle indefinido 
	//solo se podra terminar cuando se cumpla la condicion
	numero := 0
	for{
		fmt.Println("ponga el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100{
			break
		}
	}

	fmt.Println("numero correcto")

}