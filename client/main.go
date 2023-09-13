package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func testmethod(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)

}

func GetQueryMenuCategory(w http.ResponseWriter, r *http.Request) {
	// "query" 매개변수를 읽어옵니다.
	query := r.URL.Query().Get("query")

	// MySQL 연결 문자열
	dsn := "root:4613@tcp(127.0.0.1:3306)/kiosk"

	// MySQL 데이터베이스에 연결
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// SELECT 문 실행
	rows, err := db.Query("SELECT id, food_name, category FROM menu WHERE category=?", query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 결과 처리
	for rows.Next() {
		var id int
		var foodName, category string

		err := rows.Scan(&id, &foodName, &category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("이름 : %s / 카테고리 : %s\n", foodName, category)
	}
	// 클라이언트에 응답을 보냅니다.
	fmt.Fprintf(w, "클라이언트가 '%s' 요청을 보냈습니다.", query)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", testmethod)
	router.HandleFunc("/query", GetQueryMenuCategory)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
