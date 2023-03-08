package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
)

func leoArchivo(){
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil{
		fmt.Println("error al leer el archivo")
	}
	fmt.Println(string(archivo))
}

func leoArchivo2(){
	archivo, err := os.Open("./archivo.txt")
	if err != nil{
		fmt.Println("error al leer el archivo")
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan(){
		registro := scanner.Text()
		fmt.Printf("linea > "+ registro+ "\n" )
	}
}

func main(){
	leoArchivo()
	leoArchivo2()
}