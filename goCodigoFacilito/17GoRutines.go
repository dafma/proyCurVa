package main 
import (
	"fmt"
	"time"
	"strings"
	)

//ser concurrente significa dividir un problema en direrentes ejecuciones que se ejecutan de manera concurrente\al mismo tiempo
//los hilos en la computadora tiene cierto limite de hilos, pero los gorutines son infinitos


func main() {
	//para que una funcion se ejecute en un gorutine se pone go y luego la ejecucion de la funcion
	 go mi_nombre_lentoo("Marco")

	//funcion anonima 
	go func(){
		fmt.Println("Que aburridooo!")	
		var wait string
		fmt.Scanln(&wait)	
	}() 

}


func mi_nombre_lentoo(name string){
	letras := strings.Split(name, "")

	for _, letra := range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}