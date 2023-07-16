package deferpanic

import "fmt"

func VamosDefer(){
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final") // para dar una instrucción al final de todo
	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic(){
	a:=1
	if a==1{
		panic("Se encontró el valor 1")
	}
}