package main 
import (
	"net/http"
	"fmt"
	)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "<h1>Hola mundo</h1>")
	})

	mux.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "<h1>Hola prueba</h1>")
	})

	mux.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "<h1>Hola usuario</h1>")
	})

	http.ListenAndServe(":8080", mux)
}