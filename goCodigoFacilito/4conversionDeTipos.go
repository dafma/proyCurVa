package main 
import (
	"fmt"
	"strconv")


func main(){
	edad := 22
	//convertimos edad a string con Itoa
	edad_str := strconv.Itoa(edad)
	//convertir un string a un entero con Atoi
	edad_int, _ := strconv.Atoi("42")

	fmt.Println("Mi edad es"+edad_str)
	fmt.Println(edad_int+10)
}