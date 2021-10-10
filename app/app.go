package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	address := "127.0.0.1"
	port := "9090"

	router := mux.NewRouter()
	router.HandleFunc("/login", Login)
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}
