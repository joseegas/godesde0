package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func Multiplicar()string{
	scanner:=bufio.NewScanner(os.Stdin)

	for{
		fmt.Println("Ingrese número : ")
		if scanner.Scan() {
			numero, err=strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}else{
				break;
			}
		}
	}

	fmt.Println("La tabla del número",numero,"es :")

	for i:=1;i<=10;i++{
		texto+=fmt.Sprintf("%d x %d = %d\n",numero,i,numero*i)
	}

	return texto
}