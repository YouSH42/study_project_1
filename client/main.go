package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL 연결 문자열
	dsn := "root:4613@tcp(127.0.0.1:3306)/tutorialDatabase"

	// MySQL 데이터베이스에 연결
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 데이터베이스 연결 테스트
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MySQL 데이터베이스에 연결되었습니다.")

	// 여기에 데이터베이스 쿼리나 작업을 수행하는 코드를 추가할 수 있습니다.
	// SELECT 문 실행
	rows, err := db.Query("SELECT id, name, category FROM menu")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 결과 처리
	for rows.Next() {
		var id int
		var name, category string

		err := rows.Scan(&id, &name, &category)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Name: %s, Category: %s\n", id, name, category)
	}

	// 에러 처리
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
