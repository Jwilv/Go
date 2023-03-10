package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//defer se ejecuta recien cuando salga de la funcion
	archivo := "archivoFantasma.txt"
	file, err := os.Open(archivo)
	defer file.Close()
	if err != nil {
		fmt.Println("error al abrir el archivo")
		os.Exit(1)
	}
}
