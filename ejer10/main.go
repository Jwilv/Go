package main

import "fmt"

func main() {
	//creamops un map con los index de tipo string y va a tener valores tipo string
	paises := make(map[string]string)

	//setear valores
	paises["argentina"] = "buenos aires"
	paises["esp"] = "barcelona"

	fmt.Println(paises)

	//segundo ejemplo

	//aca le reservamos 5 espacios
	//paises := make(map[string]string, 5)

	//otra manera de iniciar
	//sin make
	//le decimos que los index van a ser string y que va a almacenar tipos de datos int
	//y le cargamos los valores
	campeonato := map[string]int{
		"barcelona": 59,
		"real madrid": 51,
	}
}
