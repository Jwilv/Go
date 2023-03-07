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

//genero hombre 

type hombre struct{
	edad int 
	float float32
	peso float32
	respirando bool 
	pensando bool 
	comiendo bool 
}

//genero mujer 

type mujer struct{
	edad int 
	float float32
	peso float32
	respirando bool 
	pensando bool 
	comiendo bool 
}

func main(){

}