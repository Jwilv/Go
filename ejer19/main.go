package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	//server en el puento 3000
	http.ListenAndServe(":3000", nil)
	//ruta principal
	http.HandleFunc("/", home)
}
