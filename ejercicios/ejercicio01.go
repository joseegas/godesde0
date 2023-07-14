package ejercicios

import (
	"strconv"
)

func ConvNumerico0(valor string)(int,string){
	num,_:=strconv.Atoi(valor)
	if num>100{
		return num,"es mayor a 100"
	}else{
		return num,"es menor a 100"
	}
}

func ConvNumerico(valor string)(int,string){
	num,err:=strconv.Atoi(valor)

	if err!=nil{
		return 0, "Hubo un error "+err.Error()
	}

	if num>=100{
		return num,"es igual o mayor a 100"
	}else{
		return num,"es menor a 100"
	}
}