package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"registerlogin/database"
	"registerlogin/routers"
)

func main() {
	database.InitDB()

	router := mux.NewRouter()
	routers.InitRouters(router)

	log.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
