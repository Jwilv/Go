package main

import (
	"fmt"
	// "bufio"
	// "os"
	"io/ioutil"
)

func leoArchivo(){
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil{
		fmt.Println("error al leer el archivo")
	}
	fmt.Println(string(archivo))
}

func main(){
	leoArchivo()
}