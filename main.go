package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"taxi-platcont-go/src/controller"
	"taxi-platcont-go/src/helper"
	"taxi-platcont-go/src/middleware"
	"taxi-platcont-go/src/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	controller.SessionMgr = helper.NewSessionMgr("cookie-token", 3600)
	router := mux.NewRouter().StrictSlash(true)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	middleware.EnableCORS(router)
	router.HandleFunc("/", inits)
	routes.RutasClientesCars(router)
	routes.RutasCliente(router)
	routes.RutasServicios(router)

	fmt.Printf("server listening on port %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

func inits(w http.ResponseWriter, r *http.Request) {
	datos := map[string]string{"Api": "Platcont taxi", "Version": "2.0.1", "Author": "Deybin Yoni Gil Perez"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(datos)
}
