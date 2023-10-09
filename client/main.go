package main

import (
	"log"
	"net/http"

	handler "github.com/YouSH42/study_project_1/client/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Home).Methods("GET")
	router.HandleFunc("/adddata", handler.AddDataHandler).Methods("POST")
	router.HandleFunc("/foodricelist", handler.FoodListRiceHandler).Methods("GET")
	router.HandleFunc("/foodnoodlelist", handler.FoodListNoodleHandler).Methods("GET")
	router.HandleFunc("/order", handler.OrderHandler).Methods("GET")
	router.HandleFunc("/message", handler.MessageHandler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
