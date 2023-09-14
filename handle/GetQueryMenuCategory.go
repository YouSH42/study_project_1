package handle

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

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
	rows, err := db.Query("SELECT id, food_name, price, category FROM menu WHERE category=?", query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 결과 처리
	for rows.Next() {
		var id, price int
		var foodName, category string

		err := rows.Scan(&id, &foodName, &price, &category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "이름 : %s / 가격 : %d / 카테고리 : %s\n", foodName, price, category)
	}
}
