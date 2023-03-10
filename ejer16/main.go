package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
	"os"
)

func ejemPanic(){
a:=1
if a==1{
	//aborta el programa
	panic("se encontro el valor")
}
}

func ejemDefer(){
		//defer se ejecuta recien cuando salga de la funcion
		archivo := "archivoFantasma.txt"
		file, err := os.Open(archivo)
		defer file.Close()
		if err != nil {
			fmt.Println("error al abrir el archivo")
			os.Exit(1)
		}
}

func main() {
	//ejemDefer()
	ejemPanic()
}
