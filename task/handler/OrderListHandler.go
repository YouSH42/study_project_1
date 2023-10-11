package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type OrderItem struct {
	Id       int       `json:id`
	FoodName string    `json:"name"`
	Price    int       `json:"price"`
	Category string    `json:"category"`
	Time     time.Time `json:"time"`
}

func OrderListHandler(w http.ResponseWriter, r *http.Request) {
	// 데이터베이스 연결
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName+"?parseTime=true")
	Error(err)
	defer db.Close()

	// 데이터베이스에서 데이터 가져오기
	rows, err := db.Query("SELECT * FROM orders")
	Error(err)
	defer rows.Close()

	// 데이터를 슬라이스에 저장
	var OrderData []OrderItem
	for rows.Next() {
		var item OrderItem
		if err := rows.Scan(&item.Id, &item.FoodName, &item.Price, &item.Category, &item.Time); err != nil {
			log.Fatal(err)
		}
		//fmt.Println(item.ID, item.Name, item.Price, item.Category)
		OrderData = append(OrderData, item)
	}
	// JSON 형식으로 데이터 전송
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(OrderData)
}
