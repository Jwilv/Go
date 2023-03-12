package main

import "time"

//le decimos que el argumento canal va a ser un chan de tipo time.Duration
func bucle(canal chan time.Duration){
	inicio := time.Now()
}

func main() {
	//creamos un canal con make
	//vamos a crear un chan = chanel/canal
	//de tipo time.Duration
	canal1 := make(chan time.Duration)
}