package main

import "fmt"

var state bool 

func main(){

	state = true 

	if state == true {
		fmt.Println(state);
	}else{
		fmt.Println("es falso")
	}

}