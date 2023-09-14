package main

import (
	"log"
	"net/http"

	handle "github.com/YouSH42/study_project_1/handle"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handle.Home)
	router.HandleFunc("/query", handle.GetQueryMenuCategory).Methods("GET")

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
