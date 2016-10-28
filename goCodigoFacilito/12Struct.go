package main 

import "fmt"


type User struct{
	edad int
	nombre, apellido string
	
}

func main() {
 	//fmt.Println(User{nombre:"Marco", apellido:"Romero",})
 	marco := User{nombre:"Marco", apellido:"Romero",}

 	//declaracion por orden para los struct
 	daniel :=User{12, "casa", "caso"}



 	fmt.Println(marco)
 	fmt.Println(daniel)

}