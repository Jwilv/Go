package main

import (
	"fmt"
	"strings"
	"time"
)

func nombreLento(nombre string){
letras := strings.Split(nombre, "")
for _, letra := range letras{
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(letra)
}
}

func main(){
	//se ejecuta de forma asincrona 
	//como si fuera en paralelo 
	//pero cuando termina el programa go no espera a que se termine de ejecutar 
	go nombreLento("wilvers")
	fmt.Println("estoy aca ")
	var x string
	fmt.Scanln(&x)
}