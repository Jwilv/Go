package main

//interface 
type humano interface{
	respirar()
	comer()
	pensar()
	sexo() string
}
type animal interface{
	respirar()
	comer()
	Escarnivoro() bool
}

type vegetal interface{
	ClasificacionVegetal() string 
}

func main(){

}