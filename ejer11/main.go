package main

import (
	"fmt"
	"time"

	user "./user"
)

type juan struct{
	user.Usuario
}
func main(){
	//establecemos un nuevo objeto usuario
	user := new(juan)
	user.AltaUsuario(10, "juan", time.Now(), true)
}


