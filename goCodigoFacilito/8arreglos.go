package main 

import "fmt"

func main() {
	//var arreglo[10]int
	arreglo := [3]int{1,2,3}
	fmt.Println(arreglo)

	for i :=0; i<len(arreglo); i++{
		fmt.Println(arreglo[i])
	}

	//matriz mulidimencional
	var matriz[2][2]int 
	fmt.Println(matriz)
}