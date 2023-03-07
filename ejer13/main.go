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
	EsCarnivoro() bool
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

//mundo animal 

type perro struct{
	comiendo bool 
	respirando bool 
	carnivoro bool 
}

func (this *perro) comer(){ this.comiendo = true }
func (this *perro) respirar(){this.respirando = true}
func (this *perro) EsCarnivoro()bool {return this.carnivoro }

func animalesRespirando(animal animal){
	animal.respirar()
	fmt.Println("soy un animal y estoy respirando")
}

func animalCarnivoro(animal animal) int {
	if(animal.EsCarnivoro()){
		return 1 
	}
	
	return 0
	}
	
func main(){
pedro := new(hombre)
pedro.esHombre = true 
HumanosRespirando(pedro)
maria := new(mujer)
HumanosRespirando(maria)
}