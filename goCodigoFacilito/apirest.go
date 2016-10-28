package main 
import (
	"github.com/gorilla/mux"
	"net/http"
	"log")


type Person struct{
	ID string `json:"id", omitempy"`
	Firstname string `json:"firstname", omitempy"`
	Lastname string `json:"lastname", omitempy"`
	Address *Adress `json:"id", omitempy"`
}

type Address struct{
	City string `json:"city", omitempy"`
	State string `json:"state", omitempy"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){

}
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
	
}
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){
	
}
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){
	
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID:"1", Firstname: "Nic", Lastname:"Raboy", &Address{City: "Dublin", State: "California" }})
	people = append(people, Person{ID:"2", Firstname: "Maria", Lastname:"Raboy"})
	router.HandleFunc("/people", GetPeopleEndpoint).Method("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Method("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Method("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Method("DLETE")
	log.Fatal(http.ListenAndServe(":123456", router))


}