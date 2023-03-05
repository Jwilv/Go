package main

import (
	user "./user"
	"fmt"
	"time"
)

type juan struct{
	user.Usuario
}
func main(){
	//establecemos un nuevo objeto usuario
	u := new(juan)
	u.AltaUsuario(10, "juan", time.Now(), true)
	fmt.Println(u.Usuario)
}


