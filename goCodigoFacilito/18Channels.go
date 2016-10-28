package main 
import "fmt"


//los gochanel nos permiten comunicar gorutines unas con otras 
func main() {

 channel := make(chan string)
 
 //channel nombre del canal un canal eque queramos y que tipo de informacion se enviaen el canal
 go func(channel chan string){
 	for{
 		var name string
 		fmt.Scanln(&name)
 		//estamos enviando name dentro del canal chanel
 		channel <- name
 	}
 }(channel)

//la diferencia entre recibir e enviar esque cuando enviamos la informacion aparece de lado derecho y el canal de lado izquierdo
 //cuando estamos recibiendo la informacion aparece del lado derecho y la variable que va a recibir la informacion de lado izquierdo
 msg := <- channel
 fmt.Println(msg)


}