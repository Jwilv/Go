package main

import (
	"fmt"
	"time"
)

//indicamos que va a ser un tipo de structura con el nombre usuario
type usuario struct{
	Id        int 
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func main(){
	//establecemos un nuevo objeto usuario
	user := new(usuario)
}


