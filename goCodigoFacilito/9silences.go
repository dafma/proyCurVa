package main 


import "fmt"

func main(){
	//puntero al arreglo
	//longitud del arreglo al que apunta
	//capacidad
	arreglo := [3]int{1,2,3}
	slice := arreglo[:2]
	fmt.Println(slice)
}

