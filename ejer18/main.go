package main

import (
	"fmt"
	"time"
)

//le decimos que el argumento canal va a ser un chan de tipo time.Duration
func bucle(canal chan time.Duration){
	inicio := time.Now()
	for i:=0;i<1000000000;i++{

	}
	final := time.Now()
	//le asignamos el valor al canal
	canal <- final.Sub(inicio)
}

func main() {
	//creamos un canal con make
	//vamos a crear un chan = chanel/canal
	//de tipo time.Duration
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Printf("llegue al final")
	// esto es como el await, todo se para hasta que canal1 
	//tenga gun valor y se guarde en msg
	msg := <- canal1
	fmt.Println(msg)
}