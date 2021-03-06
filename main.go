package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/The-Shy7/pokemon-api-template/routes"
	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/pokemon", routes.GetAllPokemon).Methods("GET")
	r.HandleFunc("/pokemon/type", routes.GetPokemonByType).Methods("GET")
	r.HandleFunc("/pokemon", routes.CreatePokemon).Methods("POST")
	r.HandleFunc("/pokemon", routes.UpdatePokemon).Methods("PUT")
	r.HandleFunc("/pokemon", routes.DeletePokemon).Methods("DELETE")

	fmt.Println("Server Running at 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	handleRequests()
}
