package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)
func main(){
//defer se ejecuta cuando salga de la funcion
archivo := "archivoFantasma.txt"
file, err := os.Open(archivo)
defer file.Close()
if err != nil{
	fmt.Println("error al abrir el archivo")
	os.Exit(1)
} 
}