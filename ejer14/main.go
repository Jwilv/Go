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
	archivo.Close()
}
func grabarArchivo(){
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil{
		fmt.Println("error en archivo")
		return
	}
	//con Fprint guardamos el string en una variable, en este caso en el archivo 
	fmt.Fprintln(archivo, "linea nueva")
	archivo.Close()
	fmt.Println("se grabo la linea")
	return 
}

func Append(archivo string, texto string) bool{
	arch, errorOpen := os.OpenFile(archivo, os.O_WRONLY | os.O_APPEND, 0644)

	}

func grabarArchivo2(){
fileName := "./nuevoArchivo.txt"
if Append(fileName, "\n que onda, linea agregada") == false{
	fmt.Println("error al modificar archivo")
}
}


func main(){
	leoArchivo()
	leoArchivo2()
	grabarArchivo()
	grabarArchivo2()
}