package main

import (
	"log"
	"net/http"

	"github.com/YouSH42/study_project_1/client/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Home).Methods("GET")
	router.HandleFunc("/adddata", handler.AddDataHandler).Methods("GET")
	router.HandleFunc("/foodricelist", handler.FoodListRiceHandler).Methods("GET")
	router.HandleFunc("/foodnoodlelist", handler.FoodListNoodleHandler).Methods("GET")
	router.HandleFunc("/order", handler.OrderHandler).Methods("GET")
	router.HandleFunc("/message", handler.MessageHandler)
	router.HandleFunc("/admin", handler.AdminHandler)
	// 정적 파일 서빙을 위한 디렉터리 설정
	fs := http.FileServer(http.Dir("static"))

	// /script 경로에서 정적 파일 제공
	http.Handle("/script/", http.StripPrefix("/script/", fs))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
