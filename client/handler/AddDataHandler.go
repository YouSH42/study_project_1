package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func AddDataHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	Error(err)

	var receivedFoodInfo FoodItem
	err = json.Unmarshal(body, &receivedFoodInfo)
	Error(err)
	fmt.Println("음식 이름 :", receivedFoodInfo.Name, "가격 :", receivedFoodInfo.Price, "카테고리 :", receivedFoodInfo.Category)
	// 데이터베이스 연결
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	Error(err)
	defer db.Close()
	//데이터 삽입 코드
	insertQuery := "INSERT INTO menu (food_name,price,category) VALUES (?, ?, ?)"
	_, err = db.Exec(insertQuery, receivedFoodInfo.Name, receivedFoodInfo.Price, receivedFoodInfo.Category)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("데이터 삽입 완료")

}
