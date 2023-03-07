package main

import "fmt"

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
	esHombre bool
}

func (this *hombre) respirar(){ this.respirando = true }
func (this *hombre) pensar(){ this.pensando = true }
func (this *hombre) comer(){ this.comiendo = true }
func (this *hombre) sexo() string{
	if( this.esHombre == true ){ return "hombre"}
	return "mujer"
}

//genero mujer 

type mujer struct{
	hombre
}

func HumanosRespirando(persona humano){
persona.respirar()
fmt.Printf("Soy un/a %s, y estoy respirando \n", persona.sexo())
}

func main(){
pedro := new(hombre)
pedro.esHombre = true 
HumanosRespirando(pedro)
maria := new(mujer)
HumanosRespirando(maria)
}