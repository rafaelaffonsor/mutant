package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"mutant/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/mutant", controllers.IsMutant).Methods("POST")

	router.HandleFunc("/stats", controllers.GetStats).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}