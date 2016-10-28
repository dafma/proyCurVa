package main 
import (
		 //"fmt"
		 "io"
	     "net/http"
	     "fmt"
	    )

func main() {
	http.HandleFunc("/holamundo",  func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Holaaaaaa!")
	});

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}

//ResponseWriter es una estructura que nos dice como vamos a responder a la peticion, hht.Request es un punteroahalcia la información
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Println("Hay una nueva peticion")
	io.WriteString(w, "Hola mundo")

}