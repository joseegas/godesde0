package files

import (
	"fmt"
	"godesde0/ejercicios"
	"os"
)

var filename string="./files/txt/tabla.txt"

func GrabaTabla(){
	var texto string=ejercicios.Multiplicar()
	archivo,err:=os.Create("./files/txt/tabla.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo"+err.Error())
	}
	fmt.Fprintln(archivo,texto)
}

func SumaTabla(){
	var texto string=ejercicios.Multiplicar()
	if !Append(filename,texto){
		fmt.Println("Error al concatenar el contenido")
	}
}

func Append(filen string,texto string)bool{
	arch,err:=os.OpenFile(filename,os.O_WRONLY|os.O_APPEND,0644)
	if err!=nil{
		fmt.Println("Error durante el Append "+err.Error())
		return false
	}

	_,err = arch.WriteString(texto)
	if err!=nil{
		fmt.Println("Error durante el WriteString "+err.Error())
		return false
	}
	
	arch.Close()
	return true
}

func LeoArchivo(){
	archivo,err:=os.ReadFile(filename)
	if err!=nil{
		fmt.Println("Error al leer archivo"+err.Error())
		return
	}

	fmt.Println(string(archivo))
}