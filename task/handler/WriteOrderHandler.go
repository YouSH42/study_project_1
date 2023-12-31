package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func WriteOrderHandler(conn net.Conn) {
	defer conn.Close()
	fmt.Println("클라이언트 연결 수락됨")

	// JSON 데이터 수신
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	// 수신된 JSON 데이터 파싱
	var receivedFoodItem FoodItem
	err = json.Unmarshal(buffer[:n], &receivedFoodItem)
	Error(err)

	// 파싱된 데이터 출력
	fmt.Printf("클라이언트에서 받은 JSON 데이터:\n%+v\n", receivedFoodItem)

	// 데이터베이스 연결
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	Error(err)
	defer db.Close()

	// 데이터 삽입 코드
	insertQuery := "INSERT INTO orders (food_name,price,category) VALUES (?, ?, ?)"
	_, err = db.Exec(insertQuery, receivedFoodItem.FoodName, receivedFoodItem.Price, receivedFoodItem.Category)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("데이터 삽입 완료")
}
