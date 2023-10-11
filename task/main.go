package main

import (
	"log"
	"net/http"

	"github.com/YouSH42/study_project_1/task/handler"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	router.HandleFunc("/orderlist", handler.OrderListHandler).Methods("GET")

	go handler.OrderListener()

	log.Println("Listening on :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
