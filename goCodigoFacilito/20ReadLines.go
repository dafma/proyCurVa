package main 

import (
	"fmt"
	"bufio"
	"os")

func main() {
	//abrimos un archivo
	archivo, err := os.Open("./hola.txt")

	if err != nil{
		fmt.Println("Hubo error")
	}

	scanner := bufio.NewScanner(archivo)

	var i int
	//cadaves que se itera el metodo escan se lee una linea
	for scanner.Scan(){
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}
	archivo.Close()
}