package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	//ruta principal
	http.HandleFunc("/", home)
	//server en el puento 3000
	//esto es como un bluble por eso las otras instruciones van arriba 
	http.ListenAndServe(":3000", nil)
}
