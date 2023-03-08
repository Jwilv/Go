package main

import "fmt"

//interface

type serVivo interface{
	vivo() bool
}
type humano interface {
	respirar()
	comer()
	pensar()
	sexo() string
	vivo() bool
}
type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	vivo() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

//genero hombre

type hombre struct {
	edad       int
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	viviendo bool 
}

func (this *hombre) respirar() { this.respirando = true }
func (this *hombre) pensar()   { this.pensando = true }
func (this *hombre) comer()    { this.comiendo = true }
func (this *hombre) sexo() string {
	if this.esHombre == true {
		return "hombre"
	}
	return "mujer"
}
func (this *hombre) vivo()bool{
	return true
}

//genero mujer

type mujer struct {
	hombre
}

func HumanosRespirando(persona humano) {
	persona.respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", persona.sexo())
}

//mundo animal

type perro struct {
	comiendo   bool
	respirando bool
	carnivoro  bool
	viviendo bool 
}

func (this *perro) comer()            { this.comiendo = true }
func (this *perro) respirar()         { this.respirando = true }
func (this *perro) EsCarnivoro() bool { return this.carnivoro }
func (this *perro) vivo() bool { return true }

func animalesRespirando(animal animal) {
	animal.respirar()
	fmt.Println("soy un animal y estoy respirando")
}

func animalCarnivoro(animal animal) int {
	if animal.EsCarnivoro() {
		return 1
	}

	return 0
}

func estoyVivo(ser serVivo)bool{
return ser.vivo()
}

func main() {
	pedro := new(hombre)
	pedro.esHombre = true
	HumanosRespirando(pedro)
	maria := new(mujer)
	HumanosRespirando(maria)
	rocco := new(perro)
	animalesRespirando(rocco)
	fmt.Println(animalCarnivoro(rocco))
	fmt.Println(estoyVivo(pedro))
	fmt.Println(estoyVivo(maria))
	fmt.Println(estoyVivo(rocco))
}
