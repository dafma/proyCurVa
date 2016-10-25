package main 

import (
	"fmt"
	"bufio",
	"os")

func main(){
	var edad int 
	fmt.Printf("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi edad es %d\n", edad)

	reader := bufo.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre:")
	nombre, err := reader.ReadString('\n')
	if err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println("Hola"+ nombre) 
	}

}