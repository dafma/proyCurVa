package main 

import "fmt"
 func main() {
 	/*
 	1.- punteros es una direccion de memoria
 	2.- enlugar del valor, enemos la dirección en la que esta el valor
 	3.- x y y =>
 	4.- x alteral el lugar que esta en la memoria

 	 */
 	
 	var x,y *int 
 	entero := 5

 	x=&entero
 	y=&entero
 	
 	fmt.Println(x)
 	fmt.Println(y)





 }
