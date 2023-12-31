package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func FoodListHandler(w http.ResponseWriter, _ *http.Request) {
	// 데이터베이스 연결
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	Error(err)
	defer db.Close()

	// 데이터베이스에서 데이터 가져오기
	rows, err := db.Query("SELECT * FROM menu")
	Error(err)
	defer rows.Close()

	// 데이터를 슬라이스에 저장
	var Fooddata []FoodItem
	for rows.Next() {
		var item FoodItem
		if err := rows.Scan(&item.Name, &item.Price, &item.Category); err != nil {
			log.Fatal(err)
		}
		Fooddata = append(Fooddata, item)
	}

	// JSON 형식으로 데이터 전송
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Fooddata)
}
