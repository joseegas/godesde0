package main

import (
	"fmt"
	"godesde0/variables"
)

func main(){
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado,texto:=variables.ConviertoaTexto(1555)
	fmt.Println(estado)
	fmt.Println(texto)
}