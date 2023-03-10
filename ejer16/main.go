package main

import (
	"fmt"
	"log"
	// "io/ioutil"
	// "log"
	"os"
)

func ejemPanic(){
	//funcio anonima para defer, por que defer solo ejecuta una sola cosa
	defer func(){
		//recover lee el ultimo panic
		reco := recover()
		if reco != nil{
			log.Fatalf("ocurrio un error que genero un panic \n %v",reco)
		}
	}()
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
