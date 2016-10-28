package main 

import "fmt"


type User struct{
	edad int
	nombre, apellido string
	
}

//metodo que imprima el nombre y apellidos concatenados
// la estructura antes del nombre del metodo y dentro un nombre que identifique a la estructura dentro del metodo string tipo del dato que va a retornar
func (this User) nombre_completo() string{
	return this.nombre + "" + this.apellido

}




func main() {
 	//fmt.Println(User{nombre:"Marco", apellido:"Romero",})
 	marco := User{nombre:"Marco", apellido:"Romero",}

 	//declaracion por orden para los struct
 	daniel :=User{12, "casa", "caso"}

 	carmelita := new(User)
 	carmelita.nombre = "Carmen"
 	carmelita.apellido = "Rosel"

 	fmt.Println(marco)
 	fmt.Println(daniel)
 	fmt.Println(carmelita.nombre_completo())
}