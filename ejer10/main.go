package main

import "fmt"

func main(){
	//creamops un map con los index de tipo string y va a tener valores tipo string
	paises := make(map[string]string)

	//setear valores 
	paises["argentina"] = "buenos aires"
	paises["esp"] = "barcelona"

	fmt.Println(paises)

	//segundo ejemplo 

	//aca le reservamos 5 espacios
	//paises := make(map[string]string, 5)
}