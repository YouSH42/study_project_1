package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func AddDataHandler(w http.ResponseWriter, r *http.Request) {
	//URL 쿼리 매개변수 가져오기
	queryParams := r.URL.Query()

	name := queryParams.Get("name")
	price := queryParams.Get("price")
	category := queryParams.Get("category")

	// fmt.Println("음식 이름 :", name, "가격 :", price, "카테고리 :", category)
	// 데이터베이스 연결
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	Error(err)
	defer db.Close()
	//데이터 삽입 코드
	insertQuery := "INSERT INTO menu (food_name,price,category) VALUES (?, ?, ?)"
	_, err = db.Exec(insertQuery, name, price, category)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("데이터 삽입 완료")

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
