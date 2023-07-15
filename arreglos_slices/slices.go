package arreglosslices

import "fmt"

var tablaS []int = []int{2,5,4}
var arreglo [10]int = [10]int{6,98,12,34,54,65,76,12}
func MuestroSlice(){
	fmt.Println(tablaS)

	porcion:= arreglo[3:] //Slice creado desde un vector desde posición 3
	porcion2:= arreglo[:5] //Slice creado desde un vector hasta la posición 5
	porcion3:= arreglo[6:7] //Slice desde posición 6 a 7 

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad(){
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int,0,0)
	for i:=0;i<100;i++{
		nums=append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(elementos), cap(nums))

}