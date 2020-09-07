package main

import (
	"api/bd"
	"api/controladores"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	api()
}

func api() {
	bd.NuevaConexionBD()
	defer bd.TerminarConexionBD()

	gorillaRoute := mux.NewRouter().StrictSlash(true)
	fmt.Println("Servidor iniciado en el puerto 8081")

	gorillaRoute.HandleFunc("/api/registro", controladores.PostUsuarioEndPoint).Methods("POST")
	gorillaRoute.HandleFunc("/api/login", controladores.GetLoginEndPoint).Methods("GET")

	gorillaRoute.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", gorillaRoute)
	log.Fatal(http.ListenAndServe(":8081", gorillaRoute))
}
