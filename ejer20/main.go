package main

var result int

func main() {

}

func suma(a, b int) int {
	return a + b
}

func resta(a, b int) int {
	return a - b
}

func multiplicacion(a, b int) int {
	return a * b
}

func middleware(f func(int, int) int) func(int, int) int {
	return func(a, b int)int{
		return f(a, b )
	}
}
