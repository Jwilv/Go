package main

import (
	"fmt"
	"time"

	"github.com/jwilv/tuto/user"

)

type juan struct {
	user.Usuario
}

func main() {
	//establecemos un nuevo objeto usuario
	u := new(juan)
	u.AltaUsuario(10, "juan", time.Now(), true)
	fmt.Println(u.Usuario)
}
