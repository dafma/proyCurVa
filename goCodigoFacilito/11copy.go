package main 

import "fmt"
 func main() {
 	//copia el minimo de lementos en ambs arreglos
 	slice := []int{1,2,3}
 	copia := make([]int, len(slice))

 	copy(copia, slice)

 	fmt.Println(slice)
 	fmt.Println(copia)
 }
