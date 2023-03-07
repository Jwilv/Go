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
	peso float32
	respirando bool 
	pensando bool 
	comiendo bool 
}

func (this *hombre) respirar(){ this.respirando = true }
func (this *hombre) pensar(){ this.pensando = true }
func (this *hombre) comer(){ this.comiendo = true }
func (this *hombre) sexo() string{ return "hombre" }

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